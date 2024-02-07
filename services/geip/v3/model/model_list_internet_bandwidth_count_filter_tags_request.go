package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInternetBandwidthCountFilterTagsRequest Request Object
type ListInternetBandwidthCountFilterTagsRequest struct {
	Body *ListResourcesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListInternetBandwidthCountFilterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternetBandwidthCountFilterTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInternetBandwidthCountFilterTagsRequest", string(data)}, " ")
}
