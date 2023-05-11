package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 模板的覆盖
type OverrideInfo struct {

	// 可用区
	AvailabilityZoneId string `json:"availability_zone_id"`

	// 虚拟机规格id
	FlavorId string `json:"flavor_id"`

	// 用户愿意为竞价实例每小时支付的最高价格
	SpotPrice *float64 `json:"spot_price,omitempty"`

	// 优先级，用户可以按照需要进行修改 值越小，越优先购买 取值范围：0 ~ +∞ 默认值是Integer.MAX_VALUE
	Priority *int32 `json:"priority,omitempty"`

	// 实例规格的权重。取值越高，单台实例满足计算力需求的能力越大，所需的实例数量越小。取值范围：大于0  可以根据指定实例规格的计算力和集群单节点最低计算力得出权重值。假设单节点最低计算力为8 vCPU、60GiB，则：8 vCPU、60GiB的实例规格权重可以设置为1 16 vCPU、120GiB的实例规格权重可以设置为2 默认值：1
	WeightedCapacity *float64 `json:"weighted_capacity,omitempty"`
}

func (o OverrideInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OverrideInfo struct{}"
	}

	return strings.Join([]string{"OverrideInfo", string(data)}, " ")
}
