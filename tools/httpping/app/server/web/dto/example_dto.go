package dto

import "github.com/xu-shi-fu/xwxy/tools/httpping/app/common/data/dxo"

type Example struct {
	ID dxo.ExampleID `json:"id"`

	Base

	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}
