package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDetachInternetBandwidthsGlobalEipSegmentRequestBody struct {

	// 请求列表
	GlobalEipSegments []DetachInternetBandwidth `json:"global_eip_segments"`
}

func (o BatchDetachInternetBandwidthsGlobalEipSegmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDetachInternetBandwidthsGlobalEipSegmentRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDetachInternetBandwidthsGlobalEipSegmentRequestBody", string(data)}, " ")
}
