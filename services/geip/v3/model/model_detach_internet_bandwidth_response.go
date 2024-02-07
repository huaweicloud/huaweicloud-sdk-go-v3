package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachInternetBandwidthResponse Response Object
type DetachInternetBandwidthResponse struct {

	// 本次请求的编号
	RequestId *string `json:"request_id,omitempty"`

	GlobalEip *DetachInternetBandwidthGlobalEips `json:"global_eip,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DetachInternetBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachInternetBandwidthResponse struct{}"
	}

	return strings.Join([]string{"DetachInternetBandwidthResponse", string(data)}, " ")
}
