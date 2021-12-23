package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSecurityPolicyRequest struct {
	// 自定义安全策略ID。

	SecurityPolicyId string `json:"security_policy_id"`
}

func (o ShowSecurityPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityPolicyRequest", string(data)}, " ")
}
