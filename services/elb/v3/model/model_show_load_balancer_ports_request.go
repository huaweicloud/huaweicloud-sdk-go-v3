package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLoadBalancerPortsRequest Request Object
type ShowLoadBalancerPortsRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`

	// port id。  支持多值查询，查询条件格式：*port_id=xxx&port_id=xxx*。
	PortId *[]string `json:"port_id,omitempty"`

	// ipv4 地址。  支持多值查询，查询条件格式：*ip_address=xxx&ip_address=xxx*。
	IpAddress *[]string `json:"ip_address,omitempty"`

	// ipv6 地址。  支持多值查询，查询条件格式：*ipv6_address=xxx&ipv6_address=xxx*。
	Ipv6Address *[]string `json:"ipv6_address,omitempty"`

	// port类型。  支持多值查询，查询条件格式：*type=xxx&type=xxx*。
	Type *[]string `json:"type,omitempty"`

	// 虚拟网络id。  支持多值查询，查询条件格式：*virsubnet_id=xxx&virsubnet_id=xxx*。
	VirsubnetId *[]string `json:"virsubnet_id,omitempty"`
}

func (o ShowLoadBalancerPortsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadBalancerPortsRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadBalancerPortsRequest", string(data)}, " ")
}
