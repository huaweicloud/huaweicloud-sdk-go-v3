package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInternetBandwidthFilterTagsRequest Request Object
type ListInternetBandwidthFilterTagsRequest struct {
	Limit *[]int32 `json:"limit,omitempty"`

	Offset *[]int32 `json:"offset,omitempty"`

	Body *ListResourcesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListInternetBandwidthFilterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternetBandwidthFilterTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInternetBandwidthFilterTagsRequest", string(data)}, " ")
}
