package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalConnectionBandwidthRequestBody 创建全域互联带宽实例的请求体。
type CreateGlobalConnectionBandwidthRequestBody struct {
	GlobalconnectionBandwidth *CreateGlobalConnectionBandwidth `json:"globalconnection_bandwidth"`
}

func (o CreateGlobalConnectionBandwidthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalConnectionBandwidthRequestBody struct{}"
	}

	return strings.Join([]string{"CreateGlobalConnectionBandwidthRequestBody", string(data)}, " ")
}
