package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListLoadbalancersByTagsRequest struct {
	Body *ListLoadbalancersByTagsRequestBody `json:"body,omitempty"`
}

func (o ListLoadbalancersByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadbalancersByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListLoadbalancersByTagsRequest", string(data)}, " ")
}
