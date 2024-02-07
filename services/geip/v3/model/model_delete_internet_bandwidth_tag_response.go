package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInternetBandwidthTagResponse Response Object
type DeleteInternetBandwidthTagResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInternetBandwidthTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInternetBandwidthTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteInternetBandwidthTagResponse", string(data)}, " ")
}
