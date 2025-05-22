package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtectedIpRequest Request Object
type ListProtectedIpRequest struct {

	// 开始查询的偏移量,默认值:0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量,默认值:2000
	Limit *int32 `json:"limit,omitempty"`

	// 实例id
	PackageId *string `json:"package_id,omitempty"`

	// 策略id
	PolicyId *string `json:"policy_id,omitempty"`

	// 防护ip
	Ip *string `json:"ip,omitempty"`

	// 本地标签
	Tag *string `json:"tag,omitempty"`
}

func (o ListProtectedIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectedIpRequest struct{}"
	}

	return strings.Join([]string{"ListProtectedIpRequest", string(data)}, " ")
}
