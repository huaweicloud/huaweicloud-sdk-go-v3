package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyAssignmentComplianceSummary 规则合规总结
type PolicyAssignmentComplianceSummary struct {

	// 规则合规状态
	ComplianceState *string `json:"compliance_state,omitempty"`

	PolicyAssignment *PolicyAssignment `json:"policy_assignment,omitempty"`

	Results *PolicyComplianceSummaryResults `json:"results,omitempty"`
}

func (o PolicyAssignmentComplianceSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyAssignmentComplianceSummary struct{}"
	}

	return strings.Join([]string{"PolicyAssignmentComplianceSummary", string(data)}, " ")
}
