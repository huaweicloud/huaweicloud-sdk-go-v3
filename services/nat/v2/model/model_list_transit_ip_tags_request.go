package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTransitIpTagsRequest struct {
}

func (o ListTransitIpTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitIpTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTransitIpTagsRequest", string(data)}, " ")
}
