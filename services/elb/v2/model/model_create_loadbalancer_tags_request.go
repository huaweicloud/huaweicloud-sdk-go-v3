package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLoadbalancerTagsRequest Request Object
type CreateLoadbalancerTagsRequest struct {

	// 负载均衡器ID
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *CreateLoadbalancerTagsRequestBody `json:"body,omitempty"`
}

func (o CreateLoadbalancerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerTagsRequest", string(data)}, " ")
}
