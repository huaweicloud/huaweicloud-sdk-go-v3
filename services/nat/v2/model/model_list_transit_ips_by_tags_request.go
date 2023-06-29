package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransitIpsByTagsRequest Request Object
type ListTransitIpsByTagsRequest struct {
	Body *ListTagResourceInstancesRequestBody `json:"body,omitempty"`
}

func (o ListTransitIpsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitIpsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTransitIpsByTagsRequest", string(data)}, " ")
}
