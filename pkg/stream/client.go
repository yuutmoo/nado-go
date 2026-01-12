package stream

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/coder/websocket"
	"github.com/yuutmoo/nado-go/log"
	"github.com/yuutmoo/nado-go/pkg/common"
	"github.com/yuutmoo/nado-go/pkg/signer"
	"math"
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
	log.Println("websocket client connected")
	go c.readLoop()
	go c.keepAlive()
	return nil
}

func (c *StreamClient) reconnect() {
	const (
		minDelay = 1 * time.Second
		maxDelay = 60 * time.Second
		factor   = 2.0
	)

	attempts := 0

	for {
		select {
		case <-c.ctx.Done():
			log.Debug("stream: stop reconnecting as context is cancelled")
			return
		default:
		}
		if attempts > 0 {
			delay := time.Duration(float64(minDelay) * math.Pow(factor, float64(attempts-1)))
			if delay > maxDelay {
				delay = maxDelay
			}

			jitter := time.Duration(rand.Int64n(int64(delay) / 5))
			actualDelay := delay + jitter

			log.Debugf("stream: waiting %v before next reconnect attempt (attempt %d)...", actualDelay, attempts)

			timer := time.NewTimer(actualDelay)
			select {
			case <-c.ctx.Done():
				timer.Stop()
				return
			case <-timer.C:
			}
		}

		log.Debug("stream: attempting to reconnect...")
		attempts++
		err := c.Dial(context.Background())
		if err != nil {
			log.Errorf("stream: reconnect failed: %v", err)
			continue
		}
		if c.signer != nil {
			log.Debug("stream: re-authenticating...")
			if err := c.Authenticate(c.ctx); err != nil {
				log.Errorf("stream: auth failed after reconnect: %v", err)
				continue
			}
		}

		log.Debug("stream: reconnected, restoring subscriptions...")
		var subErr error
		c.subscriptions.Range(func(key, value interface{}) bool {
			stream := value.(WSStream)
			err := c.writeJSON(c.ctx, WSRequest{
				Method: "subscribe",
				Stream: stream,
				ID:     c.requestID.Add(1),
			})

			if err != nil {
				log.Errorf("stream: resubscribe failed for %s: %v", key, err)
				subErr = err
				return false
			}
			return true
		})
		if subErr != nil {
			continue
		}
		log.Infof("stream: reconnected and subscriptions restored after %d attempts", attempts)
		attempts = 0
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
			log.Errorf("stream: read error: %v\n", err)
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
			log.Errorf("stream: control response error: %v\n", resp.Error)
		}
		if resp.Id != 0 {
			log.Errorf("stream: sub response received for id %d\n", resp.Id)
		}
	}
}

func (c *StreamClient) dispatch(streamType string, data []byte) {

	ch := c.getWorker(streamType)

	select {
	case ch <- data:
	default:
		log.Debugf("stream: [%s] processor queue full, dropping message\n", streamType)
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
					log.Errorf("stream: ping failed: %v, triggering reconnect\n", err)
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

	go func() {
		defer func() {
			c.workersMu.Lock()
			delete(c.workers, streamType)
			close(ch)
			c.workersMu.Unlock()
		}()
		c.streamProcessor(streamType, ch)
	}()

	return ch
}

func (c *StreamClient) streamProcessor(streamType string, ch chan []byte) {
	log.Printf("stream: worker processor started for [%s]\n", streamType)
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
