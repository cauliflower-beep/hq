package model

type Stock struct {
	Market    int16   // 市场id
	StockCode string  // 股票 id
	PriceNow  float64 // 现价
}
