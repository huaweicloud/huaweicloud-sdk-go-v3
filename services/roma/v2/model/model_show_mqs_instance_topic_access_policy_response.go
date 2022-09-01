package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMqsInstanceTopicAccessPolicyResponse struct {

	// topic名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 策略列表。
	Policies *[]ShowMqsInstanceTopicAccessPolicyRespPolicies `json:"policies,omitempty" xml:"policies"`

	// 权限策略的总数。
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 查询权限策略的数量。
	Size *int32 `json:"size,omitempty" xml:"size"`

	// topic名称。
	Operation      *string `json:"operation,omitempty" xml:"operation"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMqsInstanceTopicAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMqsInstanceTopicAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowMqsInstanceTopicAccessPolicyResponse", string(data)}, " ")
}
