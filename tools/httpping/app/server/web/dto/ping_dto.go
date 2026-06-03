package dto

import "github.com/xu-shi-fu/xwxy/tools/httpping/app/common/data/dxo"

type Ping struct {
	ID dxo.PingID `json:"id"`

	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
