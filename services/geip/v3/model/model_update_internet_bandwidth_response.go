package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInternetBandwidthResponse Response Object
type UpdateInternetBandwidthResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	InternetBandwidth *UpdateInternetBandwidth `json:"internet_bandwidth,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInternetBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInternetBandwidthResponse struct{}"
	}

	return strings.Join([]string{"UpdateInternetBandwidthResponse", string(data)}, " ")
}
