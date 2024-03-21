package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInternetBandwidthTagsRequest Request Object
type ShowInternetBandwidthTagsRequest struct {
	ResourceId string `json:"resource_id"`
}

func (o ShowInternetBandwidthTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInternetBandwidthTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowInternetBandwidthTagsRequest", string(data)}, " ")
}
