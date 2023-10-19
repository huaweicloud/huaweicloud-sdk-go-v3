package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AggregatePolicyComplianceSummaryResult 合规总结结果
type AggregatePolicyComplianceSummaryResult struct {
	ResourceDetails *PolicyComplianceSummaryUnit `json:"resource_details,omitempty"`

	AssignmentDetails *PolicyComplianceSummaryUnit `json:"assignment_details,omitempty"`

	// 分组名称
	GroupName *string `json:"group_name,omitempty"`

	// 帐号名称
	GroupAccountName *string `json:"group_account_name,omitempty"`
}

func (o AggregatePolicyComplianceSummaryResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregatePolicyComplianceSummaryResult struct{}"
	}

	return strings.Join([]string{"AggregatePolicyComplianceSummaryResult", string(data)}, " ")
}
