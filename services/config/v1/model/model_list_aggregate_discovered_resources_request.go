package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAggregateDiscoveredResourcesRequest Request Object
type ListAggregateDiscoveredResourcesRequest struct {

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`

	Body *AggregateDiscoveredResourcesRequest `json:"body,omitempty"`
}

func (o ListAggregateDiscoveredResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAggregateDiscoveredResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListAggregateDiscoveredResourcesRequest", string(data)}, " ")
}
