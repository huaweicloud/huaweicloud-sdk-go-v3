package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建域间带宽实例的请求体。
type CreateInterRegionBandwidthRequestBody struct {
	InterRegionBandwidth *CreateInterRegionBandwidth `json:"inter_region_bandwidth"`
}

func (o CreateInterRegionBandwidthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInterRegionBandwidthRequestBody struct{}"
	}

	return strings.Join([]string{"CreateInterRegionBandwidthRequestBody", string(data)}, " ")
}
