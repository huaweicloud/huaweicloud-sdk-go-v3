package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchAttachInternetBandwidthsGlobalEipSegmentRequestBody struct {

	// 请求列表
	GlobalEipSegments []AttachInternetBandwidth `json:"global_eip_segments"`
}

func (o BatchAttachInternetBandwidthsGlobalEipSegmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachInternetBandwidthsGlobalEipSegmentRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAttachInternetBandwidthsGlobalEipSegmentRequestBody", string(data)}, " ")
}
