package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunAggregateResourceQueryRequest Request Object
type RunAggregateResourceQueryRequest struct {

	// 资源聚合器ID。
	AggregatorId string `json:"aggregator_id"`

	Body *QueryRunRequestBody `json:"body,omitempty"`
}

func (o RunAggregateResourceQueryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunAggregateResourceQueryRequest struct{}"
	}

	return strings.Join([]string{"RunAggregateResourceQueryRequest", string(data)}, " ")
}
