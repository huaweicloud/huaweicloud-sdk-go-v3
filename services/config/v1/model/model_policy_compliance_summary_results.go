package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyComplianceSummaryResults 合规总结结果
type PolicyComplianceSummaryResults struct {
	ResourceDetails *PolicyComplianceSummaryUnit `json:"resource_details,omitempty"`

	AssignmentDetails *PolicyComplianceSummaryUnit `json:"assignment_details,omitempty"`
}

func (o PolicyComplianceSummaryResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyComplianceSummaryResults struct{}"
	}

	return strings.Join([]string{"PolicyComplianceSummaryResults", string(data)}, " ")
}
