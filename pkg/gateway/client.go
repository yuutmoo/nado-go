package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/yuutmoo/nado-go/pkg/common"
	"github.com/yuutmoo/nado-go/pkg/signer"
	"github.com/yuutmoo/nado-go/pkg/types"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type GatewayOption func(*GatewayClient)
type GatewayClient struct {
	network      common.NetworkConfig
	httpClient   *http.Client
	signer       *signer.Signer
	productCache sync.Map
}

func NewGatewayClient(opts ...GatewayOption) *GatewayClient {
	c := &GatewayClient{
		network: common.Mainnet,

		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		signer: nil,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.syncProducts()
	return c
}

func WithPrivateKey(privateKeyHex string, subAccountName string, chainId int64) GatewayOption {
	return func(c *GatewayClient) {
		s := signer.NewSigner(privateKeyHex, subAccountName, chainId)
		c.signer = s

	}
}

func WithHTTPClient(client *http.Client) GatewayOption {
	return func(c *GatewayClient) {
		c.httpClient = client
	}
}

func WithNetwork(isTestnet bool) GatewayOption {
	return func(c *GatewayClient) {
		if isTestnet {
			c.network = common.Testnet
		} else {
			c.network = common.Mainnet
		}
	}
}

func (c *GatewayClient) do(ctx context.Context, method, fullURL string, payload interface{}, result interface{}) error {
	var bodyReader io.Reader

	if payload != nil {
		jsonBytes, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("marshal payload failed: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonBytes)
	}
	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("unmarshal failed: %w, body: %s", err, string(respBody))
		}
	}
	return nil
}

func (c *GatewayClient) post(ctx context.Context, fullURL string, payload interface{}, result interface{}) error {
	return c.do(ctx, http.MethodPost, fullURL, payload, result)
}

func (c *GatewayClient) get(ctx context.Context, fullURL string, queryParams url.Values, result interface{}) error {

	if len(queryParams) > 0 {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryParams.Encode())
	}
	return c.do(ctx, http.MethodGet, fullURL, nil, result)
}

func (c *GatewayClient) buildGatewayURL(path string) string {
	return c.network.GatewayREST + path
}

// buildGatewayV2URL for V2 endpoints
func (c *GatewayClient) buildGatewayV2URL(path string) string {
	return c.network.GatewayV2 + path
}

func (c *GatewayClient) buildTriggerURL(path string) string {
	return c.network.TriggerREST + path
}

func (c *GatewayClient) syncProducts() {
	resp, err := c.GetAllProducts(context.Background())
	if err != nil {
		log.Printf("sync product fail: %v", err)
		return
	}
	for _, p := range resp.Data.SpotProducts {
		c.productCache.Store(p.ProductID, types.ProductInfo{
			ProductID:  p.ProductID,
			PriceTick:  common.X18ToFloat(p.BookInfo.PriceIncrementX18),
			AmountTick: common.X18ToFloat(p.BookInfo.SizeIncrement),
			MinSize:    common.X18ToFloat(p.BookInfo.MinSize),
		})
	}

	for _, p := range resp.Data.PerpProducts {
		c.productCache.Store(p.ProductID, types.ProductInfo{
			ProductID:  p.ProductID,
			PriceTick:  common.X18ToFloat(p.BookInfo.PriceIncrementX18),
			AmountTick: common.X18ToFloat(p.BookInfo.SizeIncrement),
			MinSize:    common.X18ToFloat(p.BookInfo.MinSize),
		})
	}
}
