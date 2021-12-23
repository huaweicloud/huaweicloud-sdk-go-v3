package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePoliciesV3Request struct {
	// 需要添加授权策略的通道名称。

	StreamName string `json:"stream_name"`

	Body *CreatePolicyRuleRequest `json:"body,omitempty"`
}

func (o CreatePoliciesV3Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoliciesV3Request struct{}"
	}

	return strings.Join([]string{"CreatePoliciesV3Request", string(data)}, " ")
}
