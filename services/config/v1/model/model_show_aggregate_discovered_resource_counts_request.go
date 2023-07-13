package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAggregateDiscoveredResourceCountsRequest Request Object
type ShowAggregateDiscoveredResourceCountsRequest struct {
	Body *AggregateDiscoveredResourceCountsRequest `json:"body,omitempty"`
}

func (o ShowAggregateDiscoveredResourceCountsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAggregateDiscoveredResourceCountsRequest struct{}"
	}

	return strings.Join([]string{"ShowAggregateDiscoveredResourceCountsRequest", string(data)}, " ")
}
