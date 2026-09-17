package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlertPolicyRequest Request Object
type ShowAlertPolicyRequest struct {

	// 租户ID
	TenantId string `json:"tenant_id"`

	// 告警策略ID
	PolicyId string `json:"policy_id"`
}

func (o ShowAlertPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowAlertPolicyRequest", string(data)}, " ")
}
