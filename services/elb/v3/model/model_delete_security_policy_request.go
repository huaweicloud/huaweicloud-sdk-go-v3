package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSecurityPolicyRequest struct {
	// 自定义安全策略的ID。

	SecurityPolicyId string `json:"security_policy_id"`
}

func (o DeleteSecurityPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityPolicyRequest", string(data)}, " ")
}
