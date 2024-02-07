package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddInternetBandwidthTagsResponse Response Object
type AddInternetBandwidthTagsResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddInternetBandwidthTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddInternetBandwidthTagsResponse struct{}"
	}

	return strings.Join([]string{"AddInternetBandwidthTagsResponse", string(data)}, " ")
}
