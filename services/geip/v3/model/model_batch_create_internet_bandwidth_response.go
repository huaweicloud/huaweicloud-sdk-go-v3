package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateInternetBandwidthResponse Response Object
type BatchCreateInternetBandwidthResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	// 创建成功响应体对象
	InternetBandwidths *[]BatchCreateInternetBandwidth `json:"internet_bandwidths,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateInternetBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateInternetBandwidthResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateInternetBandwidthResponse", string(data)}, " ")
}
