package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateLoadbalancerTagsRequest struct {
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
