package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新域间带宽实例的请求体。
type UpdateInterRegionBandwidthRequestBody struct {
	InterRegionBandwidth *UpdateInterRegionBandwidth `json:"inter_region_bandwidth"`
}

func (o UpdateInterRegionBandwidthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInterRegionBandwidthRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInterRegionBandwidthRequestBody", string(data)}, " ")
}
