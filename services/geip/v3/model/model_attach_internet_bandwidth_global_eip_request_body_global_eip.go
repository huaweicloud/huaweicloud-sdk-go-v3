package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachInternetBandwidthGlobalEipRequestBodyGlobalEip 请求参数对象
type AttachInternetBandwidthGlobalEipRequestBodyGlobalEip struct {

	// 全域公网带宽的ID
	InternetBandwidthId string `json:"internet_bandwidth_id"`
}

func (o AttachInternetBandwidthGlobalEipRequestBodyGlobalEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInternetBandwidthGlobalEipRequestBodyGlobalEip struct{}"
	}

	return strings.Join([]string{"AttachInternetBandwidthGlobalEipRequestBodyGlobalEip", string(data)}, " ")
}
