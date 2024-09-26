package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalConnectionBandwidthResponse Response Object
type CreateGlobalConnectionBandwidthResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	GlobalconnectionBandwidth *GlobalConnectionBandwidth `json:"globalconnection_bandwidth"`
	HttpStatusCode            int                        `json:"-"`
}

func (o CreateGlobalConnectionBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalConnectionBandwidthResponse struct{}"
	}

	return strings.Join([]string{"CreateGlobalConnectionBandwidthResponse", string(data)}, " ")
}
