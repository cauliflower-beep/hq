package common

import "hq/internal/model"

// getBasicHQRsp
// @Description: getBasicHQ回包
type GetBasicHQRsp struct {
	Msg  string
	Data []model.Stock
}
