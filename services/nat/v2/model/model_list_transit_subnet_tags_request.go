package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransitSubnetTagsRequest Request Object
type ListTransitSubnetTagsRequest struct {
}

func (o ListTransitSubnetTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitSubnetTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTransitSubnetTagsRequest", string(data)}, " ")
}
