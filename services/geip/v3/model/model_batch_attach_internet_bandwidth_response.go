package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAttachInternetBandwidthResponse Response Object
type BatchAttachInternetBandwidthResponse struct {
	Body *string `json:"body,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAttachInternetBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachInternetBandwidthResponse struct{}"
	}

	return strings.Join([]string{"BatchAttachInternetBandwidthResponse", string(data)}, " ")
}
