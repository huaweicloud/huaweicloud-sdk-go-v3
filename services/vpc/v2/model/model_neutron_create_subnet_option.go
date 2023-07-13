package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateSubnetOption
type NeutronCreateSubnetOption struct {

	// 功能说明：子网的名称 取值范围：0-255个字符
	Name *string `json:"name,omitempty"`

	// 功能说明：子网网段 取值范围：必须是cidr格式，只支持10.0.0.0/8,172.16.0.0/12,192.168.0.0/16三个网段内的地址，掩码长度不能大于28
	Cidr string `json:"cidr"`

	// 功能说明：子网所属网络ID
	NetworkId string `json:"network_id"`

	// 功能说明：子网网关 取值范围：子网网段中的IP地址 约束：必须是ip格式
	GatewayIp *string `json:"gateway_ip,omitempty"`

	// 功能说明：IP版本信息 取值范围：4或者6(特定局点)
	IpVersion *int32 `json:"ip_version,omitempty"`

	// 功能说明：可用的IP池，allocation_pool对象参见allocation_pool对象 例如：[ { \"start\": \"10.0.0.2\", \"end\": \"10.0.0.251\"} ]每个子网的第1个和最后4个IP地址为系统保留地址。以192.168.1.0/24为例，192.168.1.0、 192.168.1.252、192.168.1.253、192.168.1.254和192.168.1.255这些地址是系统保留地址。系统预留地址默认不在allocation_pool范围内。 约束：更新时allocation_pool范围不能包含网关和广播地址的所有IP。
	AllocationPools *[]AllocationPool `json:"allocation_pools,omitempty"`

	// 功能说明：子网关联的DNS名称服务器列表 取值范围：IP地址格式例如：\"dns_nameservers\": [\"8.xx.xx.8\",\"8.xx.xx.4\"] 默认值：不填时为空，无法使用云内网DNS功能 [内网DNS地址请参见](https://support.huaweicloud.com/dns_faq/dns_faq_002.html) [通过API获取请参见](https://support.huaweicloud.com/api-dns/dns_api_69001.html)
	DnsNameservers *[]string `json:"dns_nameservers,omitempty"`

	// 功能说明：虚拟机静态路由，参见“host_route对象”表 约束：不支持设置
	HostRoutes *[]HostRoute `json:"host_routes,omitempty"`

	// 功能说明：是否启动dhcp，false表示不提供dhcp服务的能力 约束：只支持true
	EnableDhcp *bool `json:"enable_dhcp,omitempty"`

	// 功能说明：IPv6寻址模式 取值范围：dhcpv6-stateful
	Ipv6AddressMode *string `json:"ipv6_address_mode,omitempty"`

	// 功能说明：IPv6路由广播模式 取值范围：dhcpv6-stateful
	Ipv6RaMode *string `json:"ipv6_ra_mode,omitempty"`
}

func (o NeutronCreateSubnetOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateSubnetOption struct{}"
	}

	return strings.Join([]string{"NeutronCreateSubnetOption", string(data)}, " ")
}
