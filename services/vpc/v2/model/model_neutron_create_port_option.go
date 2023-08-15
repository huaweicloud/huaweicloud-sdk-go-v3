package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreatePortOption
type NeutronCreatePortOption struct {

	// 功能说明：端口的名称 取值范围：0-255个字符
	Name *string `json:"name,omitempty"`

	// 端口所属网络ID
	NetworkId string `json:"network_id"`

	// 功能说明：管理状态 约束：只支持true
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 功能说明：端口的IP地址，参见“FixedIp对象”表 约束：device_owner为neutron: VIP_PORT时最多指定一个fixed_ip，给云服务器创建IPv6端口时，必须具备一个IPv4 subnet_id和一个IPv6 subnet_id 。
	FixedIps *[]FixedIp `json:"fixed_ips,omitempty"`

	// 功能说明：作用在该端口上的安全组的ID列表 约束：不支持更新为空
	SecurityGroups *[]string `json:"security_groups,omitempty"`

	// 功能说明：端口设备所属 取值范围：目前只支持指定\"\"和\"neutron:VIP_PORT\"；neutron:VIP_PORT表示创建的是VIP
	DeviceOwner *string `json:"device_owner,omitempty"`

	// 功能说明：扩展属性：IP/Mac对列表，详情参见“allow_address_pair对象”表 约束：IP地址不允许为 “0.0.0.0”如果allowed_address_pairs配置地址池较大的CIDR（掩码小于24位），建议为该port配置一个单独的安全组硬件SDN环境不支持ip_address属性配置为CIDR格式。
	AllowedAddressPairs *[]AllowedAddressPair `json:"allowed_address_pairs,omitempty"`

	// 功能说明：扩展属性：DHCP的扩展Option，详情参见“ExtraDhcpOpt对象”表
	ExtraDhcpOpts *[]ExtraDhcpOpt `json:"extra_dhcp_opts,omitempty"`

	// 扩展属性：提供用户设置自定义信息 【使用说明】  internal_elb字段，布尔类型，普通租户可见。只有在创建内网ELB的虚拟IP的网卡时设置为true。普通租户没有权限更改该字段，由系统维护。 举例：{\"internal_elb\": true}  disable_security_groups字段，布尔类型，普通租户可见。默认为false高性能通信场景下，允许指定为true普通租户可见。仅支持创建port和读取时指定。当前仅支持指定为true，不支持指定为false 举例：{\"disable_security_groups\"：true }， 当前仅支持指定为true，不支持指定为false，指定为true时，FWaaS功能不生效。  仅对于“华北-北京二”：udp_srvports和tcp_srvports，字段，字符串类型，默认不设置udp_srvports和tcp_srvports字段。允许指定udp_srvports和tcp_srvports字段为端口号，表示这些端口的tcp报文和udp报文可支持高并发连接，但是此类报文不受ACL和安全组规则的限制。udp_srvports和tcp_srvports字段同时支持更新操作。 − 格式： {\"tcp_srvports\": \"port1 port2 port3\", \"udp_srvports\": \"port1 port2 port3\"} 端口号之间以空格间隔，最多允许指定的端口号总共为15个，端口号范围是1到65535。 − 示例：{\"tcp_srvports\": \"80 443\", \"udp_srvports\": \"53\"} 示例表示入方向目的端口为80或者443的tcp报文可支持高并发连接。入方向目的端口为53的udp报文可支持高并发连接。但是此类报文不受ACL和安全组规则的限制。
	Bindingprofile map[string]interface{} `json:"binding:profile,omitempty"`

	// 功能说明：端口安全使能标记，如果不使能则安全组和dhcp防欺骗不生效 取值范围：启用（true）或禁用（false）
	PortSecurityEnabled *bool `json:"port_security_enabled,omitempty"`

	// 绑定的vNIC类型  - normal: 软交换
	BindingvnicType *string `json:"binding:vnic_type,omitempty"`
}

func (o NeutronCreatePortOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreatePortOption struct{}"
	}

	return strings.Join([]string{"NeutronCreatePortOption", string(data)}, " ")
}
