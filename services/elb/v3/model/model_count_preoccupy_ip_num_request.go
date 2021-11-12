package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CountPreoccupyIpNumRequest struct {
	// 功能描述：LB需要部署的AZ列表 约束：若查询创建一个LB所需预占IP数时，该参数为必选

	AvailabilityZoneId *[]string `json:"availability_zone_id,omitempty"`
	// 是否启用跨VPC后端转发

	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`
	// 负载均衡器网络类型，枚举值4，6

	IpVersion *int32 `json:"ip_version,omitempty"`
	// 七层Flavor的ID。如果欲创建7层规格的弹性负载均衡实例，则该参数为必选

	L7FlavorId *string `json:"l7_flavor_id,omitempty"`
	// 负载均衡器ID。当查询创建第一个七层监听器所需预占的ip数时，该参数为必选。

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
}

func (o CountPreoccupyIpNumRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountPreoccupyIpNumRequest struct{}"
	}

	return strings.Join([]string{"CountPreoccupyIpNumRequest", string(data)}, " ")
}
