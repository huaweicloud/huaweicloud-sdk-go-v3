package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyRequest Request Object
type ListPolicyRequest struct {

	// 开始查询的偏移量,默认值:0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量,默认值:2000
	Limit *int32 `json:"limit,omitempty"`

	// 策略名
	PolicyName *string `json:"policy_name,omitempty"`
}

func (o ListPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyRequest", string(data)}, " ")
}
