package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BccOrganizationPolicyDetail 组织策略详情
type BccOrganizationPolicyDetail struct {

	// 策略ID
	Id *string `json:"id,omitempty"`

	// 组织策略名称
	Name *string `json:"name,omitempty"`

	// 策略描述
	Description *string `json:"description,omitempty"`

	// 保护类型：备份（backup）、复制(replication)。
	OperationType *string `json:"operation_type,omitempty"`

	// 账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 策略是否开启
	PolicyEnabled *bool `json:"policy_enabled,omitempty"`

	PolicyOperationDefinition *BccOrganizationPolicyDetailPolicyOperationDefinition `json:"policy_operation_definition,omitempty"`

	PolicyTrigger *PolicyTriggerResp `json:"policy_trigger,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 账号名称
	DomainName *string `json:"domain_name,omitempty"`

	// 组织策略生效范围
	EffectiveScope *string `json:"effective_scope,omitempty"`
}

func (o BccOrganizationPolicyDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BccOrganizationPolicyDetail struct{}"
	}

	return strings.Join([]string{"BccOrganizationPolicyDetail", string(data)}, " ")
}
