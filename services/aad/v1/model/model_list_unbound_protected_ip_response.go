package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUnboundProtectedIpResponse Response Object
type ListUnboundProtectedIpResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 防护ip列表
	Ips            *[]ProtectedIpResponse `json:"ips,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListUnboundProtectedIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUnboundProtectedIpResponse struct{}"
	}

	return strings.Join([]string{"ListUnboundProtectedIpResponse", string(data)}, " ")
}
