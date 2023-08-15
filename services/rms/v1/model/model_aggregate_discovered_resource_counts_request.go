package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AggregateDiscoveredResourceCountsRequest 查询聚合器中帐号资源计数请求体。
type AggregateDiscoveredResourceCountsRequest struct {

	// 资源聚合器ID。
	AggregatorId string `json:"aggregator_id"`

	Filter *ResourceCountsFilters `json:"filter,omitempty"`

	// 用于对资源计数进行分组的键（RESOURCE_TYPE | ACCOUNT_ID）。
	GroupByKey string `json:"group_by_key"`
}

func (o AggregateDiscoveredResourceCountsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregateDiscoveredResourceCountsRequest struct{}"
	}

	return strings.Join([]string{"AggregateDiscoveredResourceCountsRequest", string(data)}, " ")
}
