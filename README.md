# nado-go


A Go implementation of the [nado.xyz](https://nado.xyz) API, featuring dual-stream isolation and automatic EIP-712 authentication.

## Usage

### Stream
```
signer := signer.NewSigner("yourprivatekey","subAccountName",common.Mainnet.ChainID)
mgr := stream.NewStreamManager(signer)
mgr.Start(context.Background())

productID := 2
// Public Stream: Market Trades
mgr.OnTrade(func(d stream.WSTradeData) { fmt.Println(d) })
mgr.SubscribeTrade(productID)

// Private Stream: Order Updates (need signer to authenticate)
mgr.OnOrderUpdate(func(d stream.WSOrderUpdate) { fmt.Println(d) })
mgr.SubscribeOrderUpdate(&productID)
```


###  Gateway (REST API)
The Gateway client handles order execution and account management.

```go
// Place a single limit order
order, err := client.PlaceOrder(ctx, &types.PlaceOrderParams{
    ProductID: 2,
    Amount:    0.0013,
    Price:     100000,
    Direction: types.DirectionShort,
})

// Batch place multiple orders
resp, err := client.PlaceOrders(ctx, []*types.PlaceOrderParams{...}, false)

// Atomic Cancel-and-Place
resp, err := client.CancelAndPlaceOrder(ctx, &placeParams, &cancelParams)

// Cancel specific orders by Digest
resp, err := client.CancelOrders(ctx, &types.CancelOrdersParam{
    Params: []types.CancelOrderParam{{ProductID: 2, Digest: "0x..."}},
    SubAccountName: "default",
})

