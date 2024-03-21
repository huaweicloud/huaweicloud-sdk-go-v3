package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateInternetBandwidthTagsResponse Response Object
type BatchCreateInternetBandwidthTagsResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateInternetBandwidthTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateInternetBandwidthTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateInternetBandwidthTagsResponse", string(data)}, " ")
}
