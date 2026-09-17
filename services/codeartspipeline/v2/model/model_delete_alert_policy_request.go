package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlertPolicyRequest Request Object
type DeleteAlertPolicyRequest struct {

	// 租户ID
	TenantId string `json:"tenant_id"`

	// 告警策略ID
	PolicyId string `json:"policy_id"`
}

func (o DeleteAlertPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlertPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlertPolicyRequest", string(data)}, " ")
}
