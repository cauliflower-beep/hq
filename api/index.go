package api

import (
	"fmt"
	"gitlab.upchinaproduct.com/taf/tafgo/taf"
	"hq/internal/model"
	"hq/pkg/common"
	"hq/pkg/jce/HQSys"
)

var _basicHqObj = `HQSys.MarketDataServer.BasicHqObj`
var _basicHQProxy = new(HQSys.BasicHq)

type Option struct {
	HqServant string
}

// Init
//  @Description: 设置主站地址，否则为默认值
//  @param op
func Init(op Option) {
	if op.HqServant != "" {
		_basicHqObj = op.HqServant
	}
}

// GetBasicHQ
//  @Description: 获取指定股票的基本行情，所属行业名称信息
//  @notice: 查询效率较低，高频访问不能用此接口，否则会崩
//  @param stocks
func GetBasicHQ(stocks []model.Stock) common.GetBasicHQRsp {
	// 构建行情数据请求
	var hStockHQReq HQSys.HStockHqReq
	//hStockHQReq.ResetDefault()

	var vStocks []HQSys.HStockUnique
	for _, stock := range stocks {
		var vStock HQSys.HStockUnique
		vStock.ShtSetcode = stock.Market
		vStock.SCode = stock.StockCode
		vStocks = append(vStocks, vStock)
	}
	var rsp common.GetBasicHQRsp

	if len(vStocks) == 0 {
		rsp.Msg = "请求股票列表为空"
		return rsp
	}

	taf.NewCommunicator().StringToProxy(_basicHqObj, _basicHQProxy)
	var hStockHQRsp HQSys.HStockHqRsp
	hStockHqRet, err := _basicHQProxy.StockHq(&hStockHQReq, &hStockHQRsp)
	if err != nil {
		rsp.Msg = fmt.Sprintf("主站接口调用失败:ret|%v^err|%v", hStockHqRet, err)
		return rsp
	}
	var data []model.Stock
	for _, stockHq := range hStockHQRsp.VStockHq {
		var stock model.Stock
		stock.Market = stockHq.ShtSetcode
		stock.StockCode = stockHq.SCode
		stock.PriceNow = stockHq.StSimHq.FNowPrice
		data = append(data, stock)
	}
	rsp.Msg = "行情请求成功"
	rsp.Data = data
	return rsp
}
