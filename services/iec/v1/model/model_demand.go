package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Demand 租户需求
type Demand struct {

	// 所属运营商。
	Operator *string `json:"operator,omitempty"`

	// 站点需要发放的资源(组)总数。  > 实际发放实例数量为count*demand_count。
	DemandCount int32 `json:"demand_count"`

	// 线路ID。 多线路场景下，将在该线路下创建弹性公网IP。 > 覆盖规则为省级/大区时不支持指定线路ID创建边缘业务。
	PoolId *string `json:"pool_id,omitempty"`

	// 带宽类型。 如果当前带宽类型下没有带宽，自动在该带宽类型下创建带宽
	BandwidthType *string `json:"bandwidth_type,omitempty"`

	// 指定IPv6线路，使用该线路下的子网分配IPv6端口。 如果该线路下没有关联启用IPv6的子网，则创建新的子网。
	PoolIdV6 *string `json:"pool_id_v6,omitempty"`

	// 使用IPv6带宽。 边缘实例是否开启IPv6公网访问能力。如果该IPv6线路没有可用的带宽，则创建新的带宽。
	Ipv6BandwidthEnable *bool `json:"ipv6_bandwidth_enable,omitempty"`

	// 带宽类型。  边缘实例开启IPv6访问公网能力后，如果当前带宽类型下没有带宽，自动在该带宽类型下创建带宽
	Ipv6BandwidthType *string `json:"ipv6_bandwidth_type,omitempty"`
}

func (o Demand) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Demand struct{}"
	}

	return strings.Join([]string{"Demand", string(data)}, " ")
}
