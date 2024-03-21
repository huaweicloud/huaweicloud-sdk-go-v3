package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddInternetBandwidthTagsRequest Request Object
type AddInternetBandwidthTagsRequest struct {
	ResourceId string `json:"resource_id"`

	Body *CreateV2TagRequestBody `json:"body,omitempty"`
}

func (o AddInternetBandwidthTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddInternetBandwidthTagsRequest struct{}"
	}

	return strings.Join([]string{"AddInternetBandwidthTagsRequest", string(data)}, " ")
}
