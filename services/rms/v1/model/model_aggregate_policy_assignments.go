package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AggregatePolicyAssignments 聚合合规规则。
type AggregatePolicyAssignments struct {

	// 合规规则ID
	PolicyAssignmentId *string `json:"policy_assignment_id,omitempty"`

	// 合规规则名称
	PolicyAssignmentName *string `json:"policy_assignment_name,omitempty"`

	Compliance *Compliance `json:"compliance,omitempty"`

	// 源帐号ID。
	AccountId *string `json:"account_id,omitempty"`
}

func (o AggregatePolicyAssignments) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregatePolicyAssignments struct{}"
	}

	return strings.Join([]string{"AggregatePolicyAssignments", string(data)}, " ")
}
