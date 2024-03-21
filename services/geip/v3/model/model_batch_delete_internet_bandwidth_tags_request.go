package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteInternetBandwidthTagsRequest Request Object
type BatchDeleteInternetBandwidthTagsRequest struct {
	ResourceId string `json:"resource_id"`

	Body *BatchDeleteV2RequestBody `json:"body,omitempty"`
}

func (o BatchDeleteInternetBandwidthTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInternetBandwidthTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteInternetBandwidthTagsRequest", string(data)}, " ")
}
