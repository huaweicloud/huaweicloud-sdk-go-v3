package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteLoadbalancerTagsRequest struct {
	// 负载均衡器ID。

	LoadbalancerId string `json:"loadbalancer_id"`

	Body *BatchDeleteLoadbalancerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteLoadbalancerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteLoadbalancerTagsRequest", string(data)}, " ")
}
