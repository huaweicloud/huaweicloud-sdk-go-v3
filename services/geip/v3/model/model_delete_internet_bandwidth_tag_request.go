package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInternetBandwidthTagRequest Request Object
type DeleteInternetBandwidthTagRequest struct {
	ResourceId string `json:"resource_id"`

	TagKey string `json:"tag_key"`
}

func (o DeleteInternetBandwidthTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInternetBandwidthTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteInternetBandwidthTagRequest", string(data)}, " ")
}
