package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDetachInternetBandwidthsGlobalEipRequestBody 批量解绑全域公网带宽请求体对象
type BatchDetachInternetBandwidthsGlobalEipRequestBody struct {

	// 批量解绑全域公网带宽请求体对象
	GlobalEips []interface{} `json:"global_eips"`
}

func (o BatchDetachInternetBandwidthsGlobalEipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDetachInternetBandwidthsGlobalEipRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDetachInternetBandwidthsGlobalEipRequestBody", string(data)}, " ")
}
