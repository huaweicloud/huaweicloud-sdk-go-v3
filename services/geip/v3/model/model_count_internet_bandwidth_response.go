package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountInternetBandwidthResponse Response Object
type CountInternetBandwidthResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	InternetBandwidths *CountInternetBandwidths `json:"internet_bandwidths,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountInternetBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountInternetBandwidthResponse struct{}"
	}

	return strings.Join([]string{"CountInternetBandwidthResponse", string(data)}, " ")
}
