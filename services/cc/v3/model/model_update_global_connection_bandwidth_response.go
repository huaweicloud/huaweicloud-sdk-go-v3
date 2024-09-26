package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalConnectionBandwidthResponse Response Object
type UpdateGlobalConnectionBandwidthResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	GlobalconnectionBandwidth *GlobalConnectionBandwidth `json:"globalconnection_bandwidth"`
	HttpStatusCode            int                        `json:"-"`
}

func (o UpdateGlobalConnectionBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalConnectionBandwidthResponse struct{}"
	}

	return strings.Join([]string{"UpdateGlobalConnectionBandwidthResponse", string(data)}, " ")
}
