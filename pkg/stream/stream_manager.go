package stream

import (
	"context"
	"fmt"
	"github.com/yuutmoo/nado-go/pkg/signer"
	"sync"
)

type StreamManager struct {
	url    string
	signer *signer.Signer

	public  *StreamClient
	private *StreamClient // (Order, Fill, Position)

	wg sync.WaitGroup

	privateMu     sync.RWMutex
	privateRoutes map[string]struct{}
}

func NewStreamManager(s *signer.Signer) *StreamManager {

	sm := &StreamManager{

		public: NewStreamClient(),
		privateRoutes: map[string]struct{}{
			"orderUpdate":    {},
			"fill":           {},
			"positionChange": {},
		},
		signer: s,
	}

	if s != nil {
		sm.private = NewStreamClient(WithSigner(sm.signer))
	}

	return sm
}

func (m *StreamManager) Start(ctx context.Context) error {
	if err := m.public.Dial(ctx); err != nil {
		return fmt.Errorf("public stream dial failed: %w", err)
	}

	if m.private != nil {
		if err := m.private.Dial(ctx); err != nil {
			return fmt.Errorf("private stream dial failed: %w", err)
		}
		if err := m.private.Authenticate(ctx); err != nil {
			return fmt.Errorf("private stream auth failed: %w", err)
		}
	}
	return nil
}

func (m *StreamManager) Close() {
	_ = m.public.Close()
	if m.private != nil {
		_ = m.private.Close()
	}
}

func (m *StreamManager) isPrivateRoute(streamType string) bool {
	m.privateMu.RLock()
	defer m.privateMu.RUnlock()
	_, ok := m.privateRoutes[streamType]
	return ok
}

func (m *StreamManager) OnTrade(cb func(WSTradeData)) {
	m.public.OnTrade(cb)
}

func (m *StreamManager) OnBookDepth(cb func(WSDepthData)) {
	m.public.OnBookDepth(cb)
}

func (m *StreamManager) OnBestBidOffer(cb func(WSBestBidOffer)) {
	m.public.OnBestBidOffer(cb)
}

func (m *StreamManager) OnCandlestick(cb func(WSLatestCandlestick)) {
	m.public.OnCandlestick(cb)
}

func (m *StreamManager) OnFundingRate(cb func(WSFundingRate)) {
	m.public.OnFundingRate(cb)
}

func (m *StreamManager) OnLiquidation(cb func(WSLiquidation)) {
	m.public.OnLiquidation(cb)
}

func (m *StreamManager) OnOrderUpdate(cb func(WSOrderUpdate)) {
	if m.private != nil {
		m.private.OnOrderUpdate(cb)
	}
}

func (m *StreamManager) OnFundingPayment(cb func(WSFundingPayment)) {
	m.public.OnFundingPayment(cb)
}

func (m *StreamManager) OnFill(cb func(WSFill)) {
	if m.private != nil {
		m.private.OnFill(cb)
	}
}

func (m *StreamManager) OnPositionChange(cb func(WSPositionChange)) {
	if m.private != nil {
		m.private.OnPositionChange(cb)
	}
}

func (m *StreamManager) SubscribeTrade(p int) error {
	return m.public.SubscribeTrade(p)
}

func (m *StreamManager) SubscribeBookDepth(p int) error {
	return m.public.SubscribeBookDepth(p)
}

func (m *StreamManager) SubscribeBestBidOffer(p int) error {
	return m.public.SubscribeBestBidOffer(p)
}

func (m *StreamManager) SubscribeCandlestick(p int, g int) error {
	return m.public.SubscribeCandlestick(p, g)
}

func (m *StreamManager) SubscribeFundingRate(p *int) error {
	return m.public.SubscribeFundingRate(p)
}

func (m *StreamManager) SubscribeLiquidation(p *int) error {
	return m.public.SubscribeLiquidation(p)
}

func (m *StreamManager) SubscribeOrderUpdate(p *int) error {
	if m.private == nil {
		return fmt.Errorf("signer required for OrderUpdate")
	}
	return m.private.SubscribeOrderUpdate(p)
}

func (m *StreamManager) SubscribeFill(p *int) error {
	if m.private == nil {
		return fmt.Errorf("signer required for Fill")
	}
	return m.private.SubscribeFill(p)
}

func (m *StreamManager) SubscribePositionChange(p *int) error {
	if m.private == nil {
		return fmt.Errorf("signer required for PositionChange")
	}
	return m.private.SubscribePositionChange(p)
}

func (m *StreamManager) SubscribeFundingPayment(p int) error {
	return m.public.SubscribeFundingPayment(p)
}
