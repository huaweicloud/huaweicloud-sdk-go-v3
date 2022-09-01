package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 合规状态
type PolicyState struct {

	// 用户ID
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 资源区域ID
	RegionId *string `json:"region_id,omitempty" xml:"region_id"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty" xml:"resource_name"`

	// 云服务名称
	ResourceProvider *string `json:"resource_provider,omitempty" xml:"resource_provider"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type"`

	// 合规状态
	ComplianceState *string `json:"compliance_state,omitempty" xml:"compliance_state"`

	// 规则ID
	PolicyAssignmentId *string `json:"policy_assignment_id,omitempty" xml:"policy_assignment_id"`

	// 规则名称
	PolicyAssignmentName *string `json:"policy_assignment_name,omitempty" xml:"policy_assignment_name"`

	// 策略ID
	PolicyDefinitionId *string `json:"policy_definition_id,omitempty" xml:"policy_definition_id"`

	// 合规状态评估时间
	EvaluationTime *string `json:"evaluation_time,omitempty" xml:"evaluation_time"`
}

func (o PolicyState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyState struct{}"
	}

	return strings.Join([]string{"PolicyState", string(data)}, " ")
}
