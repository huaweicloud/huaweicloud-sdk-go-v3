package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronSubnet
type NeutronSubnet struct {

	// 功能说明：可用的IP池，allocation_pool对象参见 allocation_pool对象 例如：[ { \"start\": \"10.0.0.2\", \"end\": \"10.0.0.251\"} ]每个子网的第1个和最后4个IP地址为系统保留地址。以192.168.1.0/24为例，192.168.1.0、 192.168.1.252、192.168.1.253、192.168.1.254和192.168.1.255这些地址是系统保留地址。系统预留地址默认不在allocation_pool范围内 约束：更新时allocation_pool范围不能包含网关和广播地址的所有IP。
	AllocationPools []AllocationPool `json:"allocation_pools"`

	// 功能说明：子网网段 取值范围：CIDR格式，只支持10.0.0.0/8,172.16.0.0/12,192.168.0.0/16三个网段内的地址，掩码长度不能大于28 约束：当ip_version=6时，该字段不支持设置
	Cidr string `json:"cidr"`

	// 功能说明：子网关联的DNS名称服务器列表 取值范围：IP地址格式例如：\"dns_nameservers\": [\"8.xx.xx.8\",\"8.xx.xx.4\"]； 默认值：不填时为空，无法使用云内网DNS功能
	DnsNameservers []string `json:"dns_nameservers"`

	// 功能说明：是否启动dhcp 取值范围：只支持true
	EnableDhcp bool `json:"enable_dhcp"`

	// 功能说明：子网网关IP  取值范围：本子网网段内的可用IP  约束：不允许和allocation_pools地址块冲突；ip_version=6时该字段不支持设置 默认值：本子网网段内第一个可用IP
	GatewayIp string `json:"gateway_ip"`

	// 虚拟机静态路由，参见“host_route对象”表
	HostRoutes []HostRoute `json:"host_routes"`

	// 子网ID
	Id string `json:"id"`

	// 功能说明：IP协议版本 取值范围：4或6(支持后)
	IpVersion int32 `json:"ip_version"`

	// 功能说明：IPv6寻址模式 取值范围：dhcpv6-stateful
	Ipv6AddressMode string `json:"ipv6_address_mode"`

	// 功能说明：IPv6路由广播模式 取值范围：dhcpv6-stateful
	Ipv6RaMode string `json:"ipv6_ra_mode"`

	// 功能说明：子网的名称 取值范围：0-255个字符
	Name string `json:"name"`

	// 所属网络的ID
	NetworkId string `json:"network_id"`

	// 项目ID
	TenantId string `json:"tenant_id"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 功能说明：资源创建UTC时间 格式：yyyy-MM-ddTHH:mm:ss
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 功能说明：资源更新UTC时间 格式：yyyy-MM-ddTHH:mm:ss
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 子网池id 【使用说明】目前IPv4不支持，IPv6支持
	SubnetpoolId *string `json:"subnetpool_id,omitempty"`
}

func (o NeutronSubnet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronSubnet struct{}"
	}

	return strings.Join([]string{"NeutronSubnet", string(data)}, " ")
}
