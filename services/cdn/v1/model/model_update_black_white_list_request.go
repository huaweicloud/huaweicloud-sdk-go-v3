package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateBlackWhiteListRequest struct {
	// 需要设置IP黑白名单的域名id。获取方法请参见查询加速域名。

	DomainId string `json:"domain_id"`

	Body *BlackWhiteListBody `json:"body,omitempty"`
}

func (o UpdateBlackWhiteListRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateBlackWhiteListRequest struct{}"
	}

	return strings.Join([]string{"UpdateBlackWhiteListRequest", string(data)}, " ")
}
