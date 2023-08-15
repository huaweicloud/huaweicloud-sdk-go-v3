package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AggregatePolicyAssignmentsRequest 聚合合规规则请求体
type AggregatePolicyAssignmentsRequest struct {

	// 资源聚合器ID
	AggregatorId string `json:"aggregator_id"`

	Filter *AggregatePolicyAssignmentsFilters `json:"filter,omitempty"`
}

func (o AggregatePolicyAssignmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregatePolicyAssignmentsRequest struct{}"
	}

	return strings.Join([]string{"AggregatePolicyAssignmentsRequest", string(data)}, " ")
}
