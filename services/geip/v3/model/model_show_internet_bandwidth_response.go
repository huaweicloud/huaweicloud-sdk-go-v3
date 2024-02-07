package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInternetBandwidthResponse Response Object
type ShowInternetBandwidthResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	InternetBandwidth *ShowInternetBandwidth `json:"internet_bandwidth,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInternetBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInternetBandwidthResponse struct{}"
	}

	return strings.Join([]string{"ShowInternetBandwidthResponse", string(data)}, " ")
}
