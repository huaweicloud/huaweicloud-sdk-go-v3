package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateInternetBandwidthTagsRequest Request Object
type BatchCreateInternetBandwidthTagsRequest struct {
	ResourceId string `json:"resource_id"`

	Body *BatchCreateV2RequestBody `json:"body,omitempty"`
}

func (o BatchCreateInternetBandwidthTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateInternetBandwidthTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateInternetBandwidthTagsRequest", string(data)}, " ")
}
