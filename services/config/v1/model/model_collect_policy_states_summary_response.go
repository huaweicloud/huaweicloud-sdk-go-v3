package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectPolicyStatesSummaryResponse Response Object
type CollectPolicyStatesSummaryResponse struct {
	Results *PolicyComplianceSummaryResults `json:"results,omitempty"`

	// 规则合规总结列表
	PolicyAssignments *[]PolicyAssignmentComplianceSummary `json:"policy_assignments,omitempty"`
	HttpStatusCode    int                                  `json:"-"`
}

func (o CollectPolicyStatesSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectPolicyStatesSummaryResponse struct{}"
	}

	return strings.Join([]string{"CollectPolicyStatesSummaryResponse", string(data)}, " ")
}
