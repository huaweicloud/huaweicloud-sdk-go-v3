package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubnetCidr struct {

	// **参数解释**： OpenStack Neutron子网的资源ID。 **取值范围**： 带“-”的标准UUID格式。
	Id string `json:"id"`

	// **参数解释**： OpenStack Neutron子网的IP版本。 **取值范围**： - 4：IPv4子网。 - 6：IPv6子网。
	IpVersion string `json:"ip_version"`

	// **参数解释**： OpenStack Neutron子网的IP网段。 **取值范围**： CIDR格式，如IPv4网段：192.168.23.0/24，IPv6网段：2420:2023:410:d5d::/64。
	Cidr string `json:"cidr"`

	// **参数解释**： OpenStack Neutron子网的网关IP地址。 **取值范围**： 不涉及。
	GatewayIp string `json:"gateway_ip"`

	// **参数解释**： OpenStack Neutron子网是否开启DHCP功能。 **取值范围**： - true：开启DHCP功能。 - false：未开启DHCP功能。
	EnableDhcp bool `json:"enable_dhcp"`
}

func (o SubnetCidr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubnetCidr struct{}"
	}

	return strings.Join([]string{"SubnetCidr", string(data)}, " ")
}
