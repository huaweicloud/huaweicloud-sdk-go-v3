package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteInternetBandwidthTagsResponse Response Object
type BatchDeleteInternetBandwidthTagsResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeleteInternetBandwidthTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInternetBandwidthTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteInternetBandwidthTagsResponse", string(data)}, " ")
}
