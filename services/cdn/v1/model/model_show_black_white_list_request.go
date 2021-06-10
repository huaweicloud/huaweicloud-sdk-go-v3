package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowBlackWhiteListRequest struct {
	// 需要查询IP黑白名单的域名id。获取方法请参见查询加速域名。

	DomainId string `json:"domain_id"`
}

func (o ShowBlackWhiteListRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBlackWhiteListRequest struct{}"
	}

	return strings.Join([]string{"ShowBlackWhiteListRequest", string(data)}, " ")
}
