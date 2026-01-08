package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLoadbalancerTagsRequest Request Object
type ListLoadbalancerTagsRequest struct {
}

func (o ListLoadbalancerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"ListLoadbalancerTagsRequest", string(data)}, " ")
}
