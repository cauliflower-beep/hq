package common

// GetBasicHQData
// @Description:
type GetBasicHQData struct {
	Stock
	PriceNow float64 // 现价
}

// GetBasicHQRsp
// @Description: getBasicHQ回包
type GetBasicHQRsp struct {
	Msg  string
	Data []GetBasicHQData
}

// KLineData
// @Description: 获取个股k线回包
type KLineData struct {
	Msg  string
	Data []HAnalyData
}

// HAnalyData
// @Description: 个股k线
type HAnalyData struct {
	SttDateTime      HDateTime       `json:"sttDateTime"`      // 时间,历史原因, 存到盘上的数据都是 *100 以后的 DWORD 型数据
	FOpen            float64         `json:"fOpen"`            // 单位开盘价
	FHigh            float64         `json:"fHigh"`            // 单位最高价
	FLow             float64         `json:"fLow"`             // 单位最低价
	FClose           float64         `json:"fClose"`           // 单位收盘价
	FAmount          float64         `json:"fAmount"`          // 单位成交金额 期货没有成交额 实际是持仓量
	LVolume          int64           `json:"lVolume"`          // 成交量
	DSettlementPrice float64         `json:"dSettlementPrice"` // 今日结算
	SttZhiShu        HTogetherZhiShu `json:"sttZhiShu"`        // 指数信息
	UiAtpVolume      uint32          `json:"uiAtpVolume"`      // 盘后成交量
	DAtpAmount       float64         `json:"dAtpAmount"`       // 盘后成交额
	UiAtpTradeNum    uint32          `json:"uiAtpTradeNum"`    // 盘后成交笔数
	FZhenfu          float32         `json:"fZhenfu"`          // 振幅 * 100%
	FTurnoverRate    float32         `json:"fTurnoverRate"`    // 换手率 * 100%
}

type HDateTime struct {
	IDate int32 `json:"iDate"` // 日期YYYYMMDD(日线)
	// 2, 3 分钟线 需要用
	ShtDay  int16 `json:"shtDay"`  // 日
	ShtTime int16 `json:"shtTime"` // 零点以来的分钟数
}

type HTogetherZhiShu struct {
	UiVolInStock uint32  // 持仓量
	FYClose      float64 // 当日结算价
	UsUp         uint16  // 上涨家数
	UsDown       uint16  // 下跌家数
}
