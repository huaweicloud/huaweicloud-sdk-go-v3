package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInternetBandwidthResponse Response Object
type DeleteInternetBandwidthResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInternetBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInternetBandwidthResponse struct{}"
	}

	return strings.Join([]string{"DeleteInternetBandwidthResponse", string(data)}, " ")
}
