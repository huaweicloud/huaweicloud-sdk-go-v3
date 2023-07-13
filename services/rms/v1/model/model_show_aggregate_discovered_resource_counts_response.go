package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAggregateDiscoveredResourceCountsResponse Response Object
type ShowAggregateDiscoveredResourceCountsResponse struct {

	// 资源计数进行分组的键。
	GroupByKey *string `json:"group_by_key,omitempty"`

	// 分组资源计数的列表。
	GroupedResourceCounts *[]GroupedResourceCount `json:"grouped_resource_counts,omitempty"`

	// 指定过滤器的资源聚合器中存在的资源总数。
	TotalDiscoveredResources *int32 `json:"total_discovered_resources,omitempty"`
	HttpStatusCode           int    `json:"-"`
}

func (o ShowAggregateDiscoveredResourceCountsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAggregateDiscoveredResourceCountsResponse struct{}"
	}

	return strings.Join([]string{"ShowAggregateDiscoveredResourceCountsResponse", string(data)}, " ")
}
