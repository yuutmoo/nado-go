package archive

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"github.com/yuutmoo/nado-go/pkg/common"
	"io"
	"net/http"
	"time"
)

type ArchiveOption func(*ArchiveClient)
type ArchiveClient struct {
	BaseURL string

	network    common.NetworkConfig
	httpClient *http.Client
}

func NewArchiveClient(opts ...ArchiveOption) *ArchiveClient {
	c := &ArchiveClient{
		network: common.Mainnet,
		BaseURL: common.Mainnet.ArchiveREST,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(c)

	}

	return c
}

func WithHTTPClient(client *http.Client) ArchiveOption {
	return func(c *ArchiveClient) {
		c.httpClient = client
	}
}

func WithNetwork(cfg common.NetworkConfig) ArchiveOption {
	return func(c *ArchiveClient) {
		c.network = cfg
	}
}

func WithTimeout(d time.Duration) ArchiveOption {
	return func(c *ArchiveClient) {
		c.httpClient.Timeout = d
	}
}

func (c *ArchiveClient) post(ctx context.Context, fullURL string, payload interface{}, result interface{}) error {
	return c.do(ctx, http.MethodPost, fullURL, payload, result)
}

func (c *ArchiveClient) do(ctx context.Context, method, fullURL string, payload interface{}, result interface{}) error {
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
	req.Header.Set("Accept-Encoding", "gzip")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()
	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return fmt.Errorf("gzip reader failed: %w", err)
	}
	respBody, err := io.ReadAll(reader)
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
