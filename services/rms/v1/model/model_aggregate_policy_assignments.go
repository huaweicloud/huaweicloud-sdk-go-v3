package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 聚合合规策略响应列表
type AggregatePolicyAssignments struct {

	// 合规规则ID
	PolicyAssignmentId *string `json:"policy_assignment_id,omitempty"`

	// 合规规则名称
	PolicyAssignmentName *string `json:"policy_assignment_name,omitempty"`

	Compliance *Compliance `json:"compliance,omitempty"`

	// 账户ID
	AccountId *string `json:"account_id,omitempty"`
}

func (o AggregatePolicyAssignments) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregatePolicyAssignments struct{}"
	}

	return strings.Join([]string{"AggregatePolicyAssignments", string(data)}, " ")
}
