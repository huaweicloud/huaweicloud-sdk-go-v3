package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConformancePackCompliance 合规规则评估结果。
type ConformancePackCompliance struct {

	// 合规规则ID。
	PolicyAssignmentId *string `json:"policy_assignment_id,omitempty"`

	// 合规规则名称。
	PolicyAssignmentName *string `json:"policy_assignment_name,omitempty"`

	// 合规规则合规结果。
	PolicyAssignmentCompliance *string `json:"policy_assignment_compliance,omitempty"`
}

func (o ConformancePackCompliance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConformancePackCompliance struct{}"
	}

	return strings.Join([]string{"ConformancePackCompliance", string(data)}, " ")
}
