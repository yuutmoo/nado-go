package stream

import (
	"context"
	"nado/pkg/common"
	"nado/pkg/signer"
	"testing"
)

func TestStreamManager(t *testing.T) {

	sm := NewStreamManager(signer.NewSigner("", "", common.Mainnet.ChainID))
	err := sm.Start(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	pid := 2
	sm.SubscribeFill(&pid)
	sm.SubscribeOrderUpdate(&pid)
	sm.SubscribePositionChange(&pid)
	sm.SubscribeCandlestick(pid, 60)
	sm.SubscribeTrade(pid)
	sm.SubscribeBookDepth(pid)
	sm.SubscribeBestBidOffer(pid)
	sm.SubscribeFundingRate(&pid)
	sm.SubscribeLiquidation(&pid)
	sm.SubscribeFundingPayment(pid)

	select {}
}
