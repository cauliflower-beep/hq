package api

import (
	"fmt"
	"gitlab.upchinaproduct.com/taf/tafgo/taf"
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
	taf.NewCommunicator().StringToProxy(_basicHqObj, _basicHQProxy)
}

// GetBasicHQ
//  @Description: 获取指定股票的基本行情，所属行业名称信息
//  @param stocks
func GetBasicHQ(stocks []common.Stock) common.GetBasicHQRsp {
	// 构建行情数据请求
	var hStockHQReq HQSys.HStockHqReq
	hStockHQReq.ResetDefault()

	var vStocks []HQSys.HStockUnique
	for _, stock := range stocks {
		var vStock HQSys.HStockUnique
		vStock.ShtSetcode = stock.Market
		vStock.SCode = stock.StockCode
		vStocks = append(vStocks, vStock)
	}
	hStockHQReq.VStock = vStocks
	hStockHQReq.EHqData = HQSys.E_STOCK_HQ_DATA_E_SHD_SIMPLE
	var rsp common.GetBasicHQRsp

	if len(vStocks) == 0 {
		rsp.Msg = "请求股票列表为空"
		return rsp
	}

	var hStockHQRsp HQSys.HStockHqRsp
	hStockHqRet, err := _basicHQProxy.StockHq(&hStockHQReq, &hStockHQRsp)
	if err != nil {
		rsp.Msg = fmt.Sprintf("主站接口调用失败:ret|%v^err|%v", hStockHqRet, err)
		return rsp
	}
	var data []common.GetBasicHQData
	for _, stockHq := range hStockHQRsp.VStockHq {
		var stock common.GetBasicHQData
		stock.Market = stockHq.ShtSetcode
		stock.StockCode = stockHq.SCode
		stock.PriceNow = stockHq.StSimHq.FNowPrice
		data = append(data, stock)
	}
	rsp.Msg = "行情请求成功"
	rsp.Data = data
	return rsp
}

// GetDayKline
//  @Description: 请求k线数据
//  @param req
//  @return common.KLineData
func GetDayKline(req common.KLineReq) common.KLineData {
	var hKLineDataReq HQSys.HKLineDataReq

	hKLineDataReq.StHeader.ShtMarket = req.Market
	hKLineDataReq.SCode = req.StockCode
	hKLineDataReq.LDate = req.LDate
	hKLineDataReq.ShtWantNum = req.ShtWantNum
	hKLineDataReq.ELineType = HQSys.HISTORY_DATA_TYPE(req.ELineType)

	var rsp common.KLineData
	if hKLineDataReq.SCode == "" {
		rsp.Msg = "请求个股代码为空"
		return rsp
	}
	var hKLineDataRsp HQSys.HKLineDataRsp
	hKlineDataRet, err := _basicHQProxy.KLineData(&hKLineDataReq, &hKLineDataRsp)
	if err != nil {
		rsp.Msg = fmt.Sprintf("主站接口调用失败:ret|%v^err|%v", hKlineDataRet, err)
		return rsp
	}

	// 构造回包data
	var hKLineData []common.HAnalyData
	for _, hAnalyData := range hKLineDataRsp.VAnalyData {
		var kLine common.HAnalyData
		kLine.SttDateTime.IDate = hAnalyData.SttDateTime.IDate
		kLine.SttDateTime.ShtDay = hAnalyData.SttDateTime.ShtTime
		kLine.SttDateTime.ShtDay = hAnalyData.SttDateTime.ShtDay

		kLine.FOpen = hAnalyData.FOpen
		kLine.FHigh = hAnalyData.FHigh
		kLine.FLow = hAnalyData.FLow
		kLine.FClose = hAnalyData.FClose
		kLine.FAmount = hAnalyData.FAmount

		kLine.LVolume = hAnalyData.LVolume
		kLine.DSettlementPrice = hAnalyData.DSettlementPrice

		kLine.SttZhiShu.UiVolInStock = hAnalyData.SttZhiShu.UiVolInStock
		kLine.SttZhiShu.FYClose = hAnalyData.SttZhiShu.FYClose
		kLine.SttZhiShu.UsUp = hAnalyData.SttZhiShu.UsUp
		kLine.SttZhiShu.UsDown = hAnalyData.SttZhiShu.UsDown

		kLine.UiAtpVolume = hAnalyData.UiAtpVolume
		kLine.DAtpAmount = hAnalyData.DAtpAmount
		kLine.UiAtpTradeNum = hAnalyData.UiAtpTradeNum
		kLine.FZhenfu = hAnalyData.FZhenfu
		kLine.FTurnoverRate = hAnalyData.FTurnoverRate

		hKLineData = append(hKLineData, kLine)
	}
	rsp.Msg = "k线数据请求成功"
	rsp.Data = hKLineData
	return rsp
}
