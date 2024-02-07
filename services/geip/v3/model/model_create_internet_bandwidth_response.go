package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInternetBandwidthResponse Response Object
type CreateInternetBandwidthResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	InternetBandwidth *CreateInternetBandwidth `json:"internet_bandwidth,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInternetBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInternetBandwidthResponse struct{}"
	}

	return strings.Join([]string{"CreateInternetBandwidthResponse", string(data)}, " ")
}
