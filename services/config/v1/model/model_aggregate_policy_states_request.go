package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AggregatePolicyStatesRequest 聚合合规评估结果请求体
type AggregatePolicyStatesRequest struct {

	// 资源聚合器ID
	AggregatorId string `json:"aggregator_id"`

	// 源帐号ID
	AccountId *string `json:"account_id,omitempty"`

	// 用于对资源计数进行分组的键（DOMAIN）。
	GroupByKey *string `json:"group_by_key,omitempty"`
}

func (o AggregatePolicyStatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregatePolicyStatesRequest struct{}"
	}

	return strings.Join([]string{"AggregatePolicyStatesRequest", string(data)}, " ")
}
