package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAttachInternetBandwidthsGlobalEipRequestBody 批量绑定全域公网带宽请求体对象
type BatchAttachInternetBandwidthsGlobalEipRequestBody struct {

	// 批量绑定全域公网带宽请求体对象
	GlobalEips []interface{} `json:"global_eips"`
}

func (o BatchAttachInternetBandwidthsGlobalEipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachInternetBandwidthsGlobalEipRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAttachInternetBandwidthsGlobalEipRequestBody", string(data)}, " ")
}
