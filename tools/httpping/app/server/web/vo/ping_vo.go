package vo

import "github.com/xu-shi-fu/xwxy/tools/httpping/app/server/web/dto"

type Pings struct {
	Base

	Items []*dto.Ping `json:"pings"`
}
