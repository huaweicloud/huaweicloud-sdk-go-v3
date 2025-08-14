package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectPolicyAssignmentsStatesSummaryResponse Response Object
type CollectPolicyAssignmentsStatesSummaryResponse struct {

	// 规则合规状态
	ComplianceState *string `json:"compliance_state,omitempty"`

	PolicyAssignment *PolicyAssignment `json:"policy_assignment,omitempty"`

	Results        *PolicyComplianceSummaryResults `json:"results,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o CollectPolicyAssignmentsStatesSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectPolicyAssignmentsStatesSummaryResponse struct{}"
	}

	return strings.Join([]string{"CollectPolicyAssignmentsStatesSummaryResponse", string(data)}, " ")
}
