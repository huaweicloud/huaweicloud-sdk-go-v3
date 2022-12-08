package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新域间带宽的详情信息。
type UpdateInterRegionBandwidth struct {

	// 域间带宽值。
	Bandwidth int32 `json:"bandwidth"`
}

func (o UpdateInterRegionBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInterRegionBandwidth struct{}"
	}

	return strings.Join([]string{"UpdateInterRegionBandwidth", string(data)}, " ")
}
