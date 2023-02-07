package test

import (
	"fmt"
	"hq/api"
	"hq/pkg/common"
	"testing"
)

func TestGetBasicHq(t *testing.T) {
	var stocks []common.Stock
	stocks = append(stocks, common.Stock{
		Market:    0,
		StockCode: "399001",
	})
	stocks = append(stocks, common.Stock{
		Market:    1,
		StockCode: "000001",
	})
	fmt.Println("请求参数股票|", stocks)
	// 设置主站地址
	var op = api.Option{
		HqServant: `HQSys.MarketDataServer.BasicHqObj@tcp -h 172.16.8.125 -t 60000 -p 8888`,
	}
	api.Init(op)

	rsp := api.GetBasicHQ(stocks)
	fmt.Println(rsp.Msg)
	fmt.Println(rsp.Data)
}
