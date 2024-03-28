package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrganizationPolicyUpdate 更新组织策略body体
type OrganizationPolicyUpdate struct {

	// 组织策略名称
	Name *string `json:"name,omitempty"`

	// 组织策略描述
	Description *string `json:"description,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 策略是否开启
	PolicyEnabled *bool `json:"policy_enabled,omitempty"`

	PolicyOperationDefinition *PolicyoOdCreate `json:"policy_operation_definition,omitempty"`

	PolicyTrigger *PolicyTriggerReq `json:"policy_trigger,omitempty"`
}

func (o OrganizationPolicyUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationPolicyUpdate struct{}"
	}

	return strings.Join([]string{"OrganizationPolicyUpdate", string(data)}, " ")
}
