package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConformancePackComplianceDetail 合规规则评估结果详情。
type ConformancePackComplianceDetail struct {

	// 合规规则ID。
	PolicyAssignmentId *string `json:"policy_assignment_id,omitempty"`

	// 合规规则名称。
	PolicyAssignmentName *string `json:"policy_assignment_name,omitempty"`

	// 评估资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 合规规则合规结果。
	ComplianceState *string `json:"compliance_state,omitempty"`

	// 资源评估时间。
	EvaluationTime *string `json:"evaluation_time,omitempty"`
}

func (o ConformancePackComplianceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConformancePackComplianceDetail struct{}"
	}

	return strings.Join([]string{"ConformancePackComplianceDetail", string(data)}, " ")
}
