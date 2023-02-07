package common

// Stock
// @Description: 股票数据
type Stock struct {
	Market    int16  // 市场id
	StockCode string // 股票 id
}

// KLineReq
// @Description: 获取个股k线入参
type KLineReq struct {
	Stock
	LDate      int64 // 日期 YYYYMMDDHHmm
	ShtWantNum int16 // 请求个数 正数往前取，负数往后取
	ELineType  int   // 请求k线数据类型
}
