package stream

import (
	"github.com/bytedance/sonic"
	"github.com/yuutmoo/nado-go/log"
)

type ResponseHandler func(data []byte)

func (c *StreamClient) OnTrade(callback func(data WSTradeData)) {
	c.SetHandler("trade", func(raw []byte) {
		var val WSTradeData
		if err := sonic.Unmarshal(raw, &val); err == nil {
			callback(val)
		} else {
			log.Errorf("stream: unmarshal trade error: %v\n", err)
		}

	})
}

func (c *StreamClient) OnFill(callback func(data WSFill)) {
	c.SetHandler("fill", func(raw []byte) {
		var val WSFill
		if err := sonic.Unmarshal(raw, &val); err == nil {
			callback(val)
		} else {
			log.Errorf("stream: unmarshal trade error: %v\n", err)
		}

	})
}

func (c *StreamClient) OnOrderUpdate(callback func(WSOrderUpdate)) {
	c.SetHandler("order_update", func(data []byte) {
		var val WSOrderUpdate
		if err := sonic.Unmarshal(data, &val); err == nil {
			callback(val)
		} else {
			log.Errorf("stream: unmarshal order update error: %v\n", err)
		}
	})
}

func (c *StreamClient) OnPositionChange(callback func(WSPositionChange)) {
	c.SetHandler("position_change", func(data []byte) {
		var val WSPositionChange
		if err := sonic.Unmarshal(data, &val); err == nil {
			callback(val)
		} else {
			log.Errorf("stream: unmarshal position change error: %v\n", err)
		}
	})
}

func (c *StreamClient) OnBookDepth(callback func(WSDepthData)) {
	c.SetHandler("book_depth", func(data []byte) {
		var val WSDepthData
		if err := sonic.Unmarshal(data, &val); err == nil {
			callback(val)
		} else {
			log.Errorf("stream: unmarshal book depth error: %v\n", err)
		}
	})
}

func (c *StreamClient) OnBestBidOffer(callback func(WSBestBidOffer)) {
	c.SetHandler("best_bid_offer", func(data []byte) {
		var val WSBestBidOffer
		if err := sonic.Unmarshal(data, &val); err == nil {
			callback(val)
		} else {
			log.Errorf("stream: unmarshal best bid offer error: %v\n", err)
		}
	})
}

func (c *StreamClient) OnCandlestick(callback func(WSLatestCandlestick)) {
	c.SetHandler("latest_candlestick", func(data []byte) {
		var val WSLatestCandlestick
		if err := sonic.Unmarshal(data, &val); err == nil {
			callback(val)
		} else {
			log.Errorf("stream: unmarshal latest candlestick error: %v\n", err)
		}
	})
}

func (c *StreamClient) OnFundingRate(callback func(WSFundingRate)) {
	c.SetHandler("funding_rate", func(data []byte) {
		var val WSFundingRate
		if err := sonic.Unmarshal(data, &val); err == nil {
			callback(val)
		} else {
			log.Errorf("stream: unmarshal funding rate error: %v\n", err)
		}
	})
}

func (c *StreamClient) OnFundingPayment(callback func(WSFundingPayment)) {
	c.SetHandler("funding_payment", func(data []byte) {
		var val WSFundingPayment
		if err := sonic.Unmarshal(data, &val); err == nil {
			callback(val)
		} else {
			log.Errorf("stream: unmarshal funding payment error: %v\n", err)
		}
	})
}

func (c *StreamClient) OnLiquidation(callback func(WSLiquidation)) {
	c.SetHandler("liquidation", func(data []byte) {
		var val WSLiquidation
		if err := sonic.Unmarshal(data, &val); err == nil {
			callback(val)
		} else {
			log.Errorf("stream: unmarshal liquidation error: %v\n", err)
		}
	})
}
