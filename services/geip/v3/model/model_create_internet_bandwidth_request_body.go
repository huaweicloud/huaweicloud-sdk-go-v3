package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInternetBandwidthRequestBody 全域公网带宽信息
type CreateInternetBandwidthRequestBody struct {
	InternetBandwidth *CreateInternetBandwidthRequestBodyInternetBandwidth `json:"internet_bandwidth"`
}

func (o CreateInternetBandwidthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInternetBandwidthRequestBody struct{}"
	}

	return strings.Join([]string{"CreateInternetBandwidthRequestBody", string(data)}, " ")
}
