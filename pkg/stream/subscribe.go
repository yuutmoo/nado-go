package stream

import (
	"context"
	"fmt"
)

type StreamOption func(*WSStream)

func WithProductID(id int) StreamOption {
	return func(s *WSStream) { s.ProductID = &id }
}

func WithGranularity(g int) StreamOption {
	return func(s *WSStream) { s.Granularity = &g }
}

func WithSubaccount(sub string) StreamOption {
	return func(s *WSStream) { s.Subaccount = &sub }
}

func (c *StreamClient) subscribe(streamType string, opts ...StreamOption) error {
	stream := WSStream{Type: streamType}
	for _, opt := range opts {
		opt(&stream)
	}
	key := c.generateKey(stream)
	req := WSRequest{
		Method: "subscribe",
		Stream: stream,
		ID:     c.requestID.Add(1),
	}
	fmt.Println("subscribing ", key)

	err := c.writeJSON(context.Background(), req)
	if err == nil {
		c.subscriptions.Store(key, stream)
	}

	return err
}

func (c *StreamClient) generateKey(stream WSStream) string {
	key := stream.Type

	if stream.Subaccount != nil {
		key = fmt.Sprintf("%s:s:%s", key, *stream.Subaccount)
	}

	if stream.ProductID != nil {
		key = fmt.Sprintf("%s:p:%d", key, *stream.ProductID)
	}

	if stream.Granularity != nil {
		key = fmt.Sprintf("%s:g:%d", key, *stream.Granularity)
	}

	return key
}

func (c *StreamClient) SubscribeTrade(productID int) error {
	return c.subscribe("trade", WithProductID(productID))
}

func (c *StreamClient) SubscribeBestBidOffer(productID int) error {
	return c.subscribe("best_bid_offer", WithProductID(productID))
}

func (c *StreamClient) SubscribeBookDepth(productID int) error {
	return c.subscribe("book_depth", WithProductID(productID))
}

func (c *StreamClient) SubscribeFundingPayment(productID int) error {
	return c.subscribe("funding_payment", WithProductID(productID))
}

// granularity: e.g., 60, 300, 3600
func (c *StreamClient) SubscribeCandlestick(productID int, granularity int) error {
	return c.subscribe("latest_candlestick",
		WithProductID(productID),
		WithGranularity(granularity),
	)
}

func (c *StreamClient) SubscribeLiquidation(productID *int) error {
	var opts []StreamOption
	if productID != nil {
		opts = append(opts, WithProductID(*productID))
	}
	return c.subscribe("liquidation", opts...)
}

func (c *StreamClient) SubscribeFundingRate(productID *int) error {
	var opts []StreamOption
	if productID != nil {
		opts = append(opts, WithProductID(*productID))
	}
	return c.subscribe("funding_rate", opts...)
}

func (c *StreamClient) SubscribeOrderUpdate(productID *int) error {
	opts := []StreamOption{WithSubaccount(c.signer.SubAccount())}
	if productID != nil {
		opts = append(opts, WithProductID(*productID))
	}
	return c.subscribe("order_update", opts...)
}

func (c *StreamClient) SubscribeFill(productID *int) error {
	opts := []StreamOption{WithSubaccount(c.signer.SubAccount())}
	if productID != nil {
		opts = append(opts, WithProductID(*productID))
	}
	return c.subscribe("fill", opts...)
}

func (c *StreamClient) SubscribePositionChange(productID *int) error {
	opts := []StreamOption{WithSubaccount(c.signer.SubAccount())}
	if productID != nil {
		opts = append(opts, WithProductID(*productID))
	}
	return c.subscribe("position_change", opts...)
}
