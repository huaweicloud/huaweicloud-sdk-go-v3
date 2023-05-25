package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAggregateComplianceByPolicyAssignmentResponse struct {

	// 聚合合规规则的列表。
	AggregatePolicyAssignments *[]AggregatePolicyAssignments `json:"aggregate_policy_assignments,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAggregateComplianceByPolicyAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAggregateComplianceByPolicyAssignmentResponse struct{}"
	}

	return strings.Join([]string{"ListAggregateComplianceByPolicyAssignmentResponse", string(data)}, " ")
}
