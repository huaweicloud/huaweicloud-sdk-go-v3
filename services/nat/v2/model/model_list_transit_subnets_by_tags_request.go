package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransitSubnetsByTagsRequest Request Object
type ListTransitSubnetsByTagsRequest struct {
	Body *ListTagResourceInstancesRequestBody `json:"body,omitempty"`
}

func (o ListTransitSubnetsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitSubnetsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTransitSubnetsByTagsRequest", string(data)}, " ")
}
