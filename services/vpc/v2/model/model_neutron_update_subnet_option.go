package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateSubnetOption
type NeutronUpdateSubnetOption struct {

	// 功能说明：dns服务器 取值范围：IP地址格式例如：\"dns_nameservers\": [\"8.xx.xx.8\",\"8.xx.xx.4\"]，最多5个 默认值：不填时为空，无法使用云内网DNS功能 [内网DNS地址请参见](https://support.huaweicloud.com/dns_faq/dns_faq_002.html) [通过API获取请参见](https://support.huaweicloud.com/api-dns/dns_api_69001.html)
	DnsNameservers *[]string `json:"dns_nameservers,omitempty"`

	// 功能说明：是否启动dhcp，false表示不提供dhcp服务的能力 约束：只支持true
	EnableDhcp *bool `json:"enable_dhcp,omitempty"`

	// 功能说明：虚拟机静态路由，参见“host_route对象”表 约束：不支持，忽略输入信息
	HostRoutes *[]HostRoute `json:"host_routes,omitempty"`

	// 子网的名称
	Name *string `json:"name,omitempty"`

	// 功能说明：可用的IP池，allocation_pool对象参见表 allocation_pool对象 例如：[ { \"start\": \"10.0.0.2\", \"end\": \"10.0.0.251\"} ]每个子网的第1个和最后3个IP地址为系统保留地址。以192.168.1.0/24为例，192.168.1.0、 192.168.1.253、192.168.1.254和192.168.1.255这些地址是系统保留地址。系统预留地址默认不在allocation_pool范围内。 约束：更新时allocation_pool范围不能包含网关和广播地址的所有IP。
	AllocationPools *[]AllocationPool `json:"allocation_pools,omitempty"`
}

func (o NeutronUpdateSubnetOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateSubnetOption struct{}"
	}

	return strings.Join([]string{"NeutronUpdateSubnetOption", string(data)}, " ")
}
