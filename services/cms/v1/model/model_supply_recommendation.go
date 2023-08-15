package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupplyRecommendation 资源供给推荐结果
type SupplyRecommendation struct {

	// 实例规格ID
	FlavorId *string `json:"flavor_id,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 可用区ID
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
