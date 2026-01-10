package gateway

import (
	"context"
	"github.com/kr/pretty"
	"github.com/yuutmoo/nado-go/pkg/common"
	"github.com/yuutmoo/nado-go/pkg/types"
	"testing"
)

func setupTestClient(t *testing.T) *GatewayClient {
	c := NewGatewayClient(WithPrivateKey("", "", common.Mainnet.ChainID))

	return c
}

func TestPlaceOrder(t *testing.T) {
	c := setupTestClient(t)

	order, err := c.PlaceOrder(context.Background(), &types.PlaceOrderParams{
		ProductID:  2,
		Amount:     0.0013,
		Price:      100000,
		AmountTick: 0.0001,
		PriceTick:  1.0,
		Direction:  types.DirectionShort,
		OrderType:  types.OrderTypeDefault,
	})
	if err != nil {
		t.Fatalf("Failed to place order: %v", err)
	}
	pretty.Println(order)
}

func TestPlaceOrders(t *testing.T) {
	c := setupTestClient(t)
	orders := make([]*types.PlaceOrderParams, 0)
	orders = append(orders, &types.PlaceOrderParams{
		ProductID:  2,
		Amount:     0.0013,
		Price:      100000,
		AmountTick: 0.0001,
		PriceTick:  1.0,
		Direction:  types.DirectionShort,
		OrderType:  types.OrderTypeDefault,
	})
	orders = append(orders, &types.PlaceOrderParams{
		ProductID:  2,
		Amount:     0.0013,
		Price:      110000,
		AmountTick: 0.0001,
		PriceTick:  1.0,
		Direction:  types.DirectionShort,
		OrderType:  types.OrderTypeDefault,
	})
	placeOrders, err := c.PlaceOrders(context.Background(), orders, false)
	if err != nil {
		t.Fatalf("Failed to place orders: %v", err)
	}
	pretty.Println(placeOrders)

}

func TestCancelAndPlaceOrder(t *testing.T) {
	c := setupTestClient(t)

	order, err := c.CancelAndPlaceOrder(context.Background(), &types.PlaceOrderParams{
		ProductID:  2,
		Price:      110000,
		Amount:     0.0013,
		AmountTick: 0.0001,
		PriceTick:  1.0,
		Direction:  types.DirectionShort,
	}, &types.CancelOrdersParam{
		Params: []types.CancelOrderParam{
			{ProductID: 2, Digest: "0xeea87b520108b5649b14cd2236c6a537aeefcf44142c1756aedc66cffddf3468"},
		},
		SubAccountName: "default",
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Log(order)
}

func TestCancelOrders(t *testing.T) {
	c := setupTestClient(t)

	orders, err := c.CancelOrders(context.Background(), &types.CancelOrdersParam{
		Params: []types.CancelOrderParam{
			{ProductID: 2, Digest: "0xfd57771d66aff5b39c8437c299d224a9b6a7ce810b680a4f66d48ff539e90d73"},
			{ProductID: 2, Digest: "0x463ca19ebc18fb5379fcce78123bf97cf3406852b5d27bb1260dada76623dbef"},
		},
		SubAccountName: "default",
	})
	if err != nil {
		t.Fatal(err)
	}
	pretty.Println(orders)
}

func TestCancelProductOrders(t *testing.T) {
	c := setupTestClient(t)
	orders, err := c.CancelProductOrders(context.Background(), []int{3})
	if err != nil {
		t.Fatal(err)
	}
	pretty.Println(orders)
}

func TestWithdrawCollateral(t *testing.T) {
	c := setupTestClient(t)

	err := c.WithdrawCollateral(context.Background(), 3, 0, 6, false)
	if err != nil {
		t.Fatal(err)
	}
}
