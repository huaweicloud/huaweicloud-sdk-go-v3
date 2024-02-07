package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachInternetBandwidthResponse Response Object
type AttachInternetBandwidthResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	GlobalEip *AttachInternetBandwidthGlobalEip `json:"global_eip,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachInternetBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInternetBandwidthResponse struct{}"
	}

	return strings.Join([]string{"AttachInternetBandwidthResponse", string(data)}, " ")
}
