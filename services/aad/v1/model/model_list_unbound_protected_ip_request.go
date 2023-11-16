package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUnboundProtectedIpRequest Request Object
type ListUnboundProtectedIpRequest struct {

	// 防护包id
	PackageId string `json:"package_id"`

	// 开始查询的偏移量,默认值:0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量,默认值:2000
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListUnboundProtectedIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUnboundProtectedIpRequest struct{}"
	}

	return strings.Join([]string{"ListUnboundProtectedIpRequest", string(data)}, " ")
}
