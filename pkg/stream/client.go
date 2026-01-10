package stream

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/coder/websocket"
	"github.com/yuutmoo/nado-go/pkg/common"
	"github.com/yuutmoo/nado-go/pkg/signer"
	"sync"
	"sync/atomic"
	"time"
)

type StreamClient struct {
	ctx    context.Context
	cancel context.CancelFunc

	network common.NetworkConfig
	conn    *websocket.Conn
	url     string

	signer *signer.Signer

	mu         sync.Mutex
	handlersMu sync.RWMutex
	requestID  atomic.Int64

	subscriptions sync.Map

	handlers map[string]func(data []byte)

	workers   map[string]chan []byte
	workersMu sync.RWMutex
}

type StreamClientOption func(*StreamClient)

func WithPrivateKey(privateKeyHex string, subAccountName string, chainId int64) StreamClientOption {
	return func(c *StreamClient) {
		s := signer.NewSigner(privateKeyHex, subAccountName, chainId)
		c.signer = s
	}
}

func WithNetworkConfig(network common.NetworkConfig) StreamClientOption {
	return func(c *StreamClient) {
		c.network = network
	}
}

func WithSigner(s *signer.Signer) StreamClientOption {
	return func(c *StreamClient) {
		c.signer = s
	}
}

func WithURL(url string) StreamClientOption {
	return func(c *StreamClient) {
		c.url = url
	}
}
func NewStreamClient(opts ...StreamClientOption) *StreamClient {
	c := &StreamClient{
		network:  common.Mainnet,
		handlers: make(map[string]func(data []byte)),
		url:      common.Mainnet.SubscriptionsWS,
		workers:  make(map[string]chan []byte),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}
func (c *StreamClient) Dial(ctx context.Context) error {
	if c.cancel != nil {
		c.cancel()
	}
	conn, _, err := websocket.Dial(ctx, c.url, &websocket.DialOptions{CompressionMode: websocket.CompressionContextTakeover})
	if err != nil {
		return err
	}

	c.conn = conn
	c.ctx, c.cancel = context.WithCancel(context.Background())
	fmt.Println("websocket client connected")
	go c.readLoop()
	go c.keepAlive()
	return nil
}

func (c *StreamClient) reconnect() {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		fmt.Println("stream: attempting to reconnect...")
		err := c.Dial(context.Background())
		if err != nil {
			time.Sleep(time.Second * 5)
			continue
		}
		if c.signer != nil {
			fmt.Println("stream: re-authenticating...")
			if err := c.Authenticate(c.ctx); err != nil {
				fmt.Printf("stream: auth failed after reconnect: %v\n", err)
				time.Sleep(time.Second * 5)
				continue
			}
		}

		fmt.Println("stream: reconnected, restoring subscriptions...")
		c.subscriptions.Range(func(key, value interface{}) bool {
			stream := value.(WSStream)
			err := c.writeJSON(c.ctx, WSRequest{
				Method: "subscribe",
				Stream: stream,
				ID:     c.requestID.Add(1),
			})

			if err != nil {
				fmt.Printf("stream: resubscribe failed for %s: %v\n", key, err)
			}
			return true
		})
		break

	}
}

func (c *StreamClient) writeJSON(ctx context.Context, v interface{}) error {
	data, err := sonic.Marshal(v)
	if err != nil {
		return fmt.Errorf("stream: sonic marshal error: %w", err)
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil {
		return fmt.Errorf("stream: cannot write to nil connection")
	}
	return c.conn.Write(ctx, websocket.MessageText, data)
}

func (c *StreamClient) Close() error {
	if c.cancel != nil {
		c.cancel()
	}
	if c.conn != nil {
		return c.conn.Close(websocket.StatusNormalClosure, "client closing")
	}
	return nil
}

func (c *StreamClient) readLoop() {
	defer func() {
		select {
		case <-c.ctx.Done():
			return
		default:
			go c.reconnect()
		}
	}()

	for {
		msgType, data, err := c.conn.Read(c.ctx)
		if err != nil {
			if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
				return
			}
			fmt.Printf("stream: read error: %v\n", err)
			return
		}

		if msgType != websocket.MessageText {
			continue
		}
		msgTypeNode, _ := sonic.Get(data, "type")
		streamType, err := msgTypeNode.String()
		if err != nil {
			c.handleControlMessage(data)
			continue
		}

		c.dispatch(streamType, data)
	}
}

func (c *StreamClient) handleControlMessage(data []byte) {
	var resp WSResponse
	if err := sonic.Unmarshal(data, &resp); err == nil {
		if resp.Error != "" {
			fmt.Printf("stream: control response error: %v\n", resp.Error)
		}
		if resp.Id != 0 {
			fmt.Printf("stream: sub response received for id %d\n", resp.Id)
		}
	}
}

func (c *StreamClient) dispatch(streamType string, data []byte) {

	ch := c.getWorker(streamType)

	select {
	case ch <- data:
	default:
		fmt.Printf("stream: [%s] processor queue full, dropping message\n", streamType)
	}

}

func (c *StreamClient) SetHandler(key string, handler ResponseHandler) {
	c.handlersMu.Lock()
	defer c.handlersMu.Unlock()
	c.handlers[key] = handler
}

func (c *StreamClient) keepAlive() {
	ticker := time.NewTicker(25 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-ticker.C:
			c.mu.Lock()
			if c.conn != nil {
				err := c.conn.Ping(c.ctx)
				if err != nil {
					fmt.Printf("stream: ping failed: %v, triggering reconnect\n", err)
					c.conn.Close(websocket.StatusAbnormalClosure, "ping timeout")
					return
				}
			}
			c.mu.Unlock()
		}
	}
}

func (c *StreamClient) getWorker(streamType string) chan []byte {
	c.workersMu.RLock()
	ch, ok := c.workers[streamType]
	c.workersMu.RUnlock()
	if ok {
		return ch
	}

	c.workersMu.Lock()
	defer c.workersMu.Unlock()

	if ch, ok = c.workers[streamType]; ok {
		return ch
	}

	ch = make(chan []byte, 1024)
	c.workers[streamType] = ch

	go c.streamProcessor(streamType, ch)

	return ch
}

func (c *StreamClient) streamProcessor(streamType string, ch chan []byte) {
	fmt.Printf("stream: worker processor started for [%s]\n", streamType)
	for {
		select {
		case <-c.ctx.Done():
			return
		case data, ok := <-ch:
			if !ok {
				return
			}
			c.handlersMu.RLock()
			handler, ok := c.handlers[streamType]
			c.handlersMu.RUnlock()

			if ok && handler != nil {
				handler(data)
			}
		}
	}
}
