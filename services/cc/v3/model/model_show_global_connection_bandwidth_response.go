package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlobalConnectionBandwidthResponse Response Object
type ShowGlobalConnectionBandwidthResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	GlobalconnectionBandwidth *GlobalConnectionBandwidth `json:"globalconnection_bandwidth"`
	HttpStatusCode            int                        `json:"-"`
}

func (o ShowGlobalConnectionBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalConnectionBandwidthResponse struct{}"
	}

	return strings.Join([]string{"ShowGlobalConnectionBandwidthResponse", string(data)}, " ")
}
