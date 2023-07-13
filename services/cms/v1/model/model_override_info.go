package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OverrideInfo 实例的详细信息
type OverrideInfo struct {

	// 可用区ID
	AvailabilityZoneId string `json:"availability_zone_id"`

	// 实例规格ID
	FlavorId string `json:"flavor_id"`

	// 用户愿意为竞价实例每小时支付的最高价格
	SpotPrice *float64 `json:"spot_price,omitempty"`

	// 优先级。数值越小，优先级越高，优先购买。 取值范围：0到Integer.MAX_VALUE 默认值是Integer.MAX_VALUE
	Priority *int32 `json:"priority,omitempty"`

	// 实例规格的权重。取值越高，单台实例满足计算力需求的能力越大，所需的实例数量越小。 取值范围：大于0 可以根据指定实例规格的计算力和集群单节点最低计算力得出权重值。 假设单节点最低计算力为8vcpu、60GB，则8vcpu、60GB的实例规格权重可设置为1，16vcpu、120GB的实例规格权重可设置为2
	WeightedCapacity *float64 `json:"weighted_capacity,omitempty"`
}

func (o OverrideInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OverrideInfo struct{}"
	}

	return strings.Join([]string{"OverrideInfo", string(data)}, " ")
}
