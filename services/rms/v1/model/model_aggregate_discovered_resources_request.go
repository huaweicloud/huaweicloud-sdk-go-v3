package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询聚合器中帐号资源计数请求体。
type AggregateDiscoveredResourcesRequest struct {

	// 资源聚合器ID。
	AggregatorId string `json:"aggregator_id"`

	Filter *ResourcesFilters `json:"filter,omitempty"`

	// 云服务类型。
	Provider *string `json:"provider,omitempty"`

	// 资源类型。
	ResourceType *string `json:"resource_type,omitempty"`
}

func (o AggregateDiscoveredResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AggregateDiscoveredResourcesRequest struct{}"
	}

	return strings.Join([]string{"AggregateDiscoveredResourcesRequest", string(data)}, " ")
}
