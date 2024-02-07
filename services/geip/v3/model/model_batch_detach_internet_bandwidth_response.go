package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDetachInternetBandwidthResponse Response Object
type BatchDetachInternetBandwidthResponse struct {
	Body *string `json:"body,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDetachInternetBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDetachInternetBandwidthResponse struct{}"
	}

	return strings.Join([]string{"BatchDetachInternetBandwidthResponse", string(data)}, " ")
}
