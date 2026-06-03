package vo

import "github.com/xu-shi-fu/xwxy/tools/httpping/app/server/web/dto"

type Examples struct {
	Base

	Items []*dto.Example `json:"examples"`
}
