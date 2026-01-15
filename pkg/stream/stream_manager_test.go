package stream

import (
	"context"
	"github.com/yuutmoo/nado-go/log"
	"github.com/yuutmoo/nado-go/pkg/common"
	"github.com/yuutmoo/nado-go/pkg/signer"
	"testing"
)

func TestStreamManager(t *testing.T) {

	sm := NewStreamManager(signer.NewSigner("", "", common.Mainnet.ChainID))
	err := sm.Start(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	pid := 2

	sm.OnOrderUpdate(func(order WSOrderUpdate) {
		log.Debug(order)
	})

	//sm.OnFill(func(fill WSFill) {
	//	log.Debug(fill)
	//})
	//sm.OnPositionChange(func(position WSPositionChange) {
	//	log.Debug(position)
	//})
	//sm.SubscribeFill(&pid)
	sm.SubscribeOrderUpdate(&pid)
	//sm.SubscribePositionChange(&pid)

	select {}
}
