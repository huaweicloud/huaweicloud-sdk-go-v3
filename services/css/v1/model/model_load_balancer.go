package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LoadBalancer struct {

	// 负载均衡器ID。
	Id *string `json:"id,omitempty"`

	// 负载均衡器名称。
	Name *string `json:"name,omitempty"`

	// 是否独享型负载均衡器。
	Guaranteed *string `json:"guaranteed,omitempty"`

	// 资源账单信息。
	BillingInfo *string `json:"billing_info,omitempty"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// 负载均衡器所属VPC ID，即虚拟私有云ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 负载均衡器的配置状态。
	ProvisioningStatus *string `json:"provisioning_status,omitempty"`

	// 关联的listener列表。
	Listeners *[]IdListWrapper `json:"listeners,omitempty"`

	// 负载均衡器的IPv4虚拟IP地址。
	VipAddress *string `json:"vip_address,omitempty"`

	// 负载均衡器的IPv4对应的port ID。
	VipPortId *string `json:"vip_port_id,omitempty"`

	// 负载均衡器的IPv6地址。
	Ipv6VipAddress *string `json:"ipv6_vip_address,omitempty"`

	// 负载均衡器绑定的公网IP。
	Publicips *[]PublicIpInfo `json:"publicips,omitempty"`
}

func (o LoadBalancer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancer struct{}"
	}

	return strings.Join([]string{"LoadBalancer", string(data)}, " ")
}
