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

	order, err := c.PlaceOrder(context.Background(), &types.PlaceOrderParam{
		ProductID: 2,
		Amount:    0.0013,
		Price:     100000,
		Direction: types.DirectionShort,
		OrderType: types.OrderTypeDefault,
	})
	if err != nil {
		t.Fatalf("Failed to place order: %v", err)
	}
	pretty.Println(order)
}

func TestPlaceOrders(t *testing.T) {
	c := setupTestClient(t)
	orders := make([]*types.PlaceOrderParam, 0)
	orders = append(orders, &types.PlaceOrderParam{
		ProductID: 2,
		Amount:    0.0013,
		Price:     100000,
		Direction: types.DirectionShort,
		OrderType: types.OrderTypeDefault,
	})
	orders = append(orders, &types.PlaceOrderParam{
		ProductID: 2,
		Amount:    0.0013,
		Price:     110000,
		Direction: types.DirectionShort,
		OrderType: types.OrderTypeDefault,
	})
	placeOrders, err := c.PlaceOrders(context.Background(), orders, false)
	if err != nil {
		t.Fatalf("Failed to place orders: %v", err)
	}
	pretty.Println(placeOrders)

}

func TestCancelAndPlaceOrder(t *testing.T) {
	c := setupTestClient(t)

	order, err := c.CancelAndPlaceOrder(context.Background(), &types.PlaceOrderParam{
		ProductID: 2,
		Price:     110000,
		Amount:    0.0013,
		Direction: types.DirectionShort,
	}, &types.CancelOrdersParam{
		Params: []types.CancelOrderParam{
			{ProductID: 2, Digest: "0x52a62de58a2b0521ed352d7ac358aa7f88ec1138a31999c0092cbf80c188c37b"},
		},
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
			{ProductID: 2, Digest: "0x7a26be95e07ec83459a3940bd95560093abf559eadc9186fcd2473279b480b30"},
		},
	}, true)
	if err != nil {
		t.Fatal(err)
	}
	pretty.Println(orders)
}

func TestCancelProductOrders(t *testing.T) {
	c := setupTestClient(t)
	orders, err := c.CancelProductOrders(context.Background(), []int{3}, false)
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
