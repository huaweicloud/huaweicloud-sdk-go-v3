package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源供给推荐结果
type SupplyRecommendation struct {

	// 弹性云服务器规格
	FlavorId *string `json:"flavor_id,omitempty"`

	// 地域id
	RegionId *string `json:"region_id,omitempty"`

	// 可用区id
	AvailabilityZoneId *string `json:"availability_zone_id,omitempty"`

	// 推荐分数
	Score *int32 `json:"score,omitempty"`
}

func (o SupplyRecommendation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupplyRecommendation struct{}"
	}

	return strings.Join([]string{"SupplyRecommendation", string(data)}, " ")
}
