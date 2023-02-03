package test

import (
	"fmt"
	"hq/api"
	"hq/internal/model"
	"testing"
)

func TestGetBasicHq(t *testing.T) {
	var stocks []model.Stock
	stocks = append(stocks, model.Stock{
		Market:    0,
		StockCode: "399001",
	})
	stocks = append(stocks, model.Stock{
		Market:    1,
		StockCode: "000001",
	})
	// 设置主站地址
	var op = api.Option{
		HqServant: `HQSys.MarketDataServer.BasicHqObj@tcp -h 172.29.13.12 -t 60000 -p 10050`,
	}
	api.Init(op)

	rsp := api.GetBasicHQ(stocks)
	fmt.Println(rsp.Msg)
	fmt.Println(rsp.Data)
}
