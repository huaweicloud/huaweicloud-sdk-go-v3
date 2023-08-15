package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyState 合规状态
type PolicyState struct {

	// 用户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 资源区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 云服务名称
	ResourceProvider *string `json:"resource_provider,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 触发器类型，可选值：resource, period
	TriggerType *string `json:"trigger_type,omitempty"`

	// 合规状态
	ComplianceState *string `json:"compliance_state,omitempty"`

	// 规则ID
	PolicyAssignmentId *string `json:"policy_assignment_id,omitempty"`

	// 规则名称
	PolicyAssignmentName *string `json:"policy_assignment_name,omitempty"`

	// 策略ID
	PolicyDefinitionId *string `json:"policy_definition_id,omitempty"`

	// 合规状态评估时间
	EvaluationTime *string `json:"evaluation_time,omitempty"`
}

func (o PolicyState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyState struct{}"
	}

	return strings.Join([]string{"PolicyState", string(data)}, " ")
}
