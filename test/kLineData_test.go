package test

import (
	"fmt"
	"hq/api"
	"hq/pkg/common"
	"testing"
)

func TestGetDayKline(t *testing.T) {
	// 设置主站地址
	var op = api.Option{
		HqServant: `HQSys.MarketDataServer.BasicHqObj@tcp -h 172.16.8.125 -t 60000 -p 8888`,
	}
	api.Init(op)
	var req common.KLineReq
	req.Market = 0
	req.StockCode = "399001"
	req.ELineType = 0
	// 1.获取2.6这一天的k线
	req.LDate = 202302070000
	req.ShtWantNum = -1
	rsp := api.GetDayKline(req)
	fmt.Println(rsp.Msg)
	fmt.Println(rsp.Data)
	// 2.获取2.5~2.7这三天的k线
	req.LDate = 0 // 0只会向前取，包含当天的
	req.ShtWantNum = 3
	rsp = api.GetDayKline(req)
	fmt.Println(rsp.Msg)
	fmt.Println(rsp.Data)
	// 3.获取当天的k线
	req.LDate = 0
	req.ShtWantNum = 1
	rsp = api.GetDayKline(req)
	fmt.Println(rsp.Msg)
	fmt.Println(rsp.Data)
	// 4.获取2.2~2.3这两天的k线
	req.LDate = 202302010000
	req.ShtWantNum = 2
	rsp = api.GetDayKline(req)
	fmt.Println(rsp.Msg)
	fmt.Println(rsp.Data)
	// 5.获取所有的k险数据
	req.LDate = 0
	req.ShtWantNum = -1
	rsp = api.GetDayKline(req)
	fmt.Println(rsp.Msg)
	fmt.Println(rsp.Data)
}
