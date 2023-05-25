package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询指定聚合合规规则详情请求体。
type AggregatePolicyAssignmentDetailRequest struct {

	// 资源聚合器ID
	AggregatorId string `json:"aggregator_id"`

	// 租户ID
	AccountId string `json:"account_id"`

	// 合规规则ID
	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o AggregatePolicyAssignmentDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregatePolicyAssignmentDetailRequest struct{}"
	}

	return strings.Join([]string{"AggregatePolicyAssignmentDetailRequest", string(data)}, " ")
}
