package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateInternetBandwidthRequestBody struct {
	InternetBandwidth *BatchCreateInternetBandwidthRequestBodyInternetBandwidth `json:"internet_bandwidth"`
}

func (o BatchCreateInternetBandwidthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateInternetBandwidthRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateInternetBandwidthRequestBody", string(data)}, " ")
}
