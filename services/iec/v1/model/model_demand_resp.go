package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DemandResp 租户需求
type DemandResp struct {

	// 站点需要发放的资源(组)总数。  > 实际发放实例数量为count*demand_count。
	DemandCount int32 `json:"demand_count"`

	// 指定IPv6线路，使用该线路下的子网分配IPv6端口。 如果该线路下没有关联启用IPv6的子网，则创建新的子网。
	PoolIdV6 *string `json:"pool_id_v6,omitempty"`

	// 使用IPv6带宽。 边缘实例是否开启IPv6公网访问能力。如果该IPv6线路下没有带宽，则创建新的带宽。
	Ipv6BandwidthEnable *bool `json:"ipv6_bandwidth_enable,omitempty"`
}

func (o DemandResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DemandResp struct{}"
	}

	return strings.Join([]string{"DemandResp", string(data)}, " ")
}
