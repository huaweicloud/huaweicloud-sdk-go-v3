package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateGlobalConnectionBandwidthRequestBody struct {
	GlobalconnectionBandwidth *UpdateGlobalConnectionBandwidth `json:"globalconnection_bandwidth"`
}

func (o UpdateGlobalConnectionBandwidthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalConnectionBandwidthRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateGlobalConnectionBandwidthRequestBody", string(data)}, " ")
}
