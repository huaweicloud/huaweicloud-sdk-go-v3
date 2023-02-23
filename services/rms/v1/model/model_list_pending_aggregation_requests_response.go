package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPendingAggregationRequestsResponse struct {

	// 挂起的聚合请求列表。
	PendingAggregationRequests *[]PendingAggregationRequest `json:"pending_aggregation_requests,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPendingAggregationRequestsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPendingAggregationRequestsResponse struct{}"
	}

	return strings.Join([]string{"ListPendingAggregationRequestsResponse", string(data)}, " ")
}
