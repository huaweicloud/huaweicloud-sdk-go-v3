package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMqsInstanceTopicAccessPolicyResponse struct {

	// topic名称。
	Name *string `json:"name,omitempty"`

	// 策略列表。
	Policies *[]ShowMqsInstanceTopicAccessPolicyRespPolicies `json:"policies,omitempty"`

	// 权限策略的总数。
	Total *int32 `json:"total,omitempty"`

	// 查询权限策略的数量。
	Size *int32 `json:"size,omitempty"`

	// topic名称。
	Operation      *string `json:"operation,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMqsInstanceTopicAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMqsInstanceTopicAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowMqsInstanceTopicAccessPolicyResponse", string(data)}, " ")
}
