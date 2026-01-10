package stream

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/coder/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/yuutmoo/nado-go/pkg/common"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func setupMockWS(t *testing.T, handler func(conn *websocket.Conn)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, nil)
		if err != nil {
			return
		}
		defer c.Close(websocket.StatusInternalError, "the end")
		handler(c)
	}))
}

func TestStream_ReconnectAndResubscribe(t *testing.T) {
	subCount := 0
	server := setupMockWS(t, func(conn *websocket.Conn) {
		for {
			_, data, err := conn.Read(context.Background())
			if err != nil {
				break
			}
			if strings.Contains(string(data), "subscribe") {
				subCount++
			}
		}
	})
	defer server.Close()

	client := NewStreamClient(WithURL(strings.Replace(server.URL, "http", "ws", 1)))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := client.Dial(ctx)
	assert.NoError(t, err)
	client.SubscribeTrade(1)
	time.Sleep(100 * time.Millisecond)

	client.conn.Close(websocket.StatusAbnormalClosure, "simulated crash")

	time.Sleep(2 * time.Second)

	assert.GreaterOrEqual(t, subCount, 2)

}

func TestStream_DataFlowAndInterruption(t *testing.T) {
	var (
		receivedCount atomic.Int64
		stopServer    = make(chan struct{})
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close(websocket.StatusInternalError, "end")

		go func() {
			for {
				_, _, err := conn.Read(context.Background())
				if err != nil {
					return
				}
			}
		}()

		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-stopServer:
				return
			case <-ticker.C:
				tradeMsg := WSTradeData{
					Type:      "trade",
					ProductId: 1,
					Price:     "27000",
					TakerQty:  "1.5",
					Timestamp: "1700000000000",
				}
				payload, _ := sonic.Marshal(tradeMsg)
				_ = conn.Write(context.Background(), websocket.MessageText, payload)
			}
		}
	}))
	defer server.Close()

	client := NewStreamClient(WithURL(strings.Replace(server.URL, "http", "ws", 1)))

	client.OnTrade(func(data WSTradeData) {
		receivedCount.Add(1)
		fmt.Printf("Received data #%d, price: %s\n", receivedCount.Load(), data.Price)
	})

	err := client.Dial(context.Background())
	assert.NoError(t, err)

	err = client.SubscribeTrade(1)
	assert.NoError(t, err)

	time.Sleep(600 * time.Millisecond)
	countBefore := receivedCount.Load()
	assert.Greater(t, countBefore, int64(0), "should have received a message")

	fmt.Println("!!! Simulating network failure now...")
	client.mu.Lock()
	client.conn.Close(websocket.StatusAbnormalClosure, "network crash")
	client.mu.Unlock()

	fmt.Println("Waiting for client to reconnect and resume data flow...")

	assert.Eventually(t, func() bool {
		return receivedCount.Load() > countBefore
	}, 10*time.Second, 500*time.Millisecond, "should have received a message after the reconnect")

	fmt.Printf("Test Passed! Total received: %d (Before: %d)\n", receivedCount.Load(), countBefore)
	close(stopServer)
}

func TestStream_Robustness_Reconnect(t *testing.T) {
	var (
		receivedCount atomic.Int64
		stopServer    = make(chan struct{})
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close(websocket.StatusInternalError, "server closed")

		go func() {
			for {
				_, _, err := conn.Read(context.Background())
				if err != nil {
					return
				}
			}
		}()

		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-stopServer:
				return
			case <-ticker.C:
				msg := WSTradeData{
					Type:      "trade",
					ProductId: 1,
					Price:     "27000.5",
					TakerQty:  "0.1",
				}
				payload, _ := sonic.Marshal(msg)
				_ = conn.Write(context.Background(), websocket.MessageText, payload)
			}
		}
	}))
	defer server.Close()

	wsURL := strings.Replace(server.URL, "http", "ws", 1)
	client := NewStreamClient(WithURL(wsURL))

	err := client.Dial(context.Background())
	assert.NoError(t, err)

	client.OnTrade(func(data WSTradeData) {
		receivedCount.Add(1)
	})

	err = client.SubscribeTrade(1)
	assert.NoError(t, err)

	time.Sleep(500 * time.Millisecond)
	countBefore := receivedCount.Load()
	assert.Greater(t, countBefore, int64(0), "should have received a message before the disconnect")
	fmt.Printf("Step 1: Received %d messages before disconnect\n", countBefore)

	fmt.Println("Step 2: Simulating network crash (closing connection)...")
	client.mu.Lock()
	client.conn.Close(websocket.StatusAbnormalClosure, "manual crash")
	client.mu.Unlock()

	fmt.Println("Step 3: Waiting for auto-reconnect and data recovery...")

	assert.Eventually(t, func() bool {
		newCount := receivedCount.Load()
		return newCount > countBefore
	}, 10*time.Second, 500*time.Millisecond, "should have received a message after the reconnect")

	fmt.Printf("Step 4: Success! Final count: %d\n", receivedCount.Load())
	close(stopServer)
}

func TestStream_AllSubscriptions(t *testing.T) {

	client := NewStreamClient()
	_ = client.Dial(context.Background())
	counts := make(map[string]*atomic.Int64)
	streamTypes := []string{
		"trade", "book_depth", "best_bid_offer",
		"funding_rate", "liquidation", "latest_candlestick",
		"fill", "order_update", "position_change",
	}
	for _, st := range streamTypes {
		counts[st] = &atomic.Int64{}
	}

	pid := 2
	gran := 60

	client.OnTrade(func(d WSTradeData) { counts["trade"].Add(1) })
	client.OnBookDepth(func(d WSDepthData) { counts["book_depth"].Add(1) })
	client.OnBestBidOffer(func(d WSBestBidOffer) { counts["best_bid_offer"].Add(1) })
	client.OnFundingRate(func(d WSFundingRate) { counts["funding_rate"].Add(1) })
	client.OnLiquidation(func(d WSLiquidation) { counts["liquidation"].Add(1) })
	client.OnCandlestick(func(d WSLatestCandlestick) { counts["latest_candlestick"].Add(1) })

	fmt.Println("sending subscription...")
	client.SubscribeTrade(pid)
	client.SubscribeBookDepth(pid)
	client.SubscribeBestBidOffer(pid)
	client.SubscribeFundingRate(nil)
	client.SubscribeLiquidation(nil)
	client.SubscribeCandlestick(pid, gran)

	start := time.Now()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for time.Since(start) < 30*time.Second {
		<-ticker.C
		fmt.Println("-------------------------------------------")
		fmt.Printf("time: %s\n", time.Now().Format("15:04:05"))
		for _, st := range streamTypes {
			fmt.Printf("Stream [%-18s]: received %d message\n", st, counts[st].Load())
		}
	}

	assert.Greater(t, counts["trade"].Load(), int64(0), "Trade stream should receive messages")
	assert.Greater(t, counts["book_depth"].Load(), int64(0), "Depth stream should receive messages")

}

func TestStream_Authenticate(t *testing.T) {
	client := NewStreamClient(WithPrivateKey("", "default", common.Mainnet.ChainID))
	_ = client.Dial(context.Background())
	//err := client.Authenticate(context.Background())
	//assert.NoError(t, err)

	//err = client.SubscribeOrderUpdate(nil)
	//if err != nil {
	//	t.Fatal(err)
	//}
	pid := 2
	client.SubscribeTrade(pid)
	client.SubscribeBookDepth(pid)
	client.SubscribeBestBidOffer(pid)
	client.SubscribeFundingRate(nil)
	client.SubscribeLiquidation(nil)
	client.SubscribeCandlestick(pid, 60)
	client.SubscribeFundingPayment(pid)
	client.SubscribeFill(&pid)
	client.SubscribePositionChange(&pid)
	select {}
}
