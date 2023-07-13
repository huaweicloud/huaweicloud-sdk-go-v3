package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronPort
type NeutronPort struct {

	// 功能说明：端口管理状态 约束：目前支持true
	AdminStateUp bool `json:"admin_state_up"`

	// 功能说明：扩展属性：IP/Mac对列表，详情参见“allow_address_pair对象”表 约束：IP地址不允许为 “0.0.0.0”如果allowed_address_pairs配置地址池较大的CIDR（掩码小于24位），建议为该port配置一个单独的安全组硬件SDN环境不支持ip_address属性配置为CIDR格式。
	AllowedAddressPairs []AllowedAddressPair `json:"allowed_address_pairs"`

	// 扩展属性：提供用户设置自定义信息 【使用说明】  internal_elb字段，布尔类型，普通租户可见。只有在创建内网ELB的虚拟IP的网卡时设置为true。普通租户没有权限更改该字段，由系统维护。 举例：{\"internal_elb\": true}  disable_security_groups字段，布尔类型，普通租户可见。默认为false高性能通信场景下，允许指定为true普通租户可见。仅支持创建port和读取时指定。当前仅支持指定为true，不支持指定为false 举例：{\"disable_security_groups\"：true }， 当前仅支持指定为true，不支持指定为false，指定为true时，FWaaS功能不生效。  仅对于“华北-北京二”：udp_srvports和tcp_srvports，字段，字符串类型，默认不设置udp_srvports和tcp_srvports字段。允许指定udp_srvports和tcp_srvports字段为端口号，表示这些端口的tcp报文和udp报文可支持高并发连接，但是此类报文不受ACL和安全组规则的限制。udp_srvports和tcp_srvports字段同时支持更新操作。 − 格式： {\"tcp_srvports\": \"port1 port2 port3\", \"udp_srvports\": \"port1 port2 port3\"} 端口号之间以空格间隔，最多允许指定的端口号总共为15个，端口号范围是1到65535。 − 示例：{\"tcp_srvports\": \"80 443\", \"udp_srvports\": \"53\"} 示例表示入方向目的端口为80或者443的tcp报文可支持高并发连接。入方向目的端口为53的udp报文可支持高并发连接。但是此类报文不受ACL和安全组规则的限制。
	Bindingprofile *interface{} `json:"binding:profile"`

	BindingvifDetails *BindingVifDetails `json:"binding:vif_details"`

	// 功能说明：绑定的vNIC类型  - normal：软交换
	BindingvnicType string `json:"binding:vnic_type"`

	// 功能说明：端口设备ID 约束：不支持设置和更新，由系统自动维护，该字段非空的端口不允许删除。
	DeviceId string `json:"device_id"`

	// 功能说明：端口设备所属（DHCP/Router/ Nova等） 约束：不支持更新，只允许用户在创建虚拟IP端口时，为虚拟IP端口设置device_owner为neutron:VIP_PORT，当端口的该字段不为空时，仅支持该字段为neutron:VIP_PORT时的端口删除。该字段非空的端口不允许删除。
	DeviceOwner string `json:"device_owner"`

	// 功能说明：扩展属性：主网卡默认内网域名信息 约束：不支持设置和更新，由系统自动维护  - hostname：与端口dns_name一致  - ip_address：端口ipv4私有地址  - fqdn：为端口创建默认内网fqdn
	DnsAssignment []DnsAssignMent `json:"dns_assignment"`

	// 功能说明：扩展属性：主网卡默认内网DNS名称 约束：不支持设置和更新，由系统自动维护,访问该默认内网域名前，请确保子网使用当前系统提供的DNS
	DnsName string `json:"dns_name"`

	// 功能说明：扩展属性：DHCP的扩展Option，详情参见“ExtraDhcpOpt对象”表
	ExtraDhcpOpts []ExtraDhcpOpt `json:"extra_dhcp_opts"`

	// 功能说明：端口的IP地址，参见“FixedIp对象”表 约束：device_owner为neutron: VIP_PORT时最多指定一个fixed_ip，给云服务器创建IPv6端口时，必须具备一个IPv4 subnet_id和一个IPv6 subnet_id 。
	FixedIps []FixedIp `json:"fixed_ips"`

	// 端口ID
	Id string `json:"id"`

	// 功能说明：端口mac地址 约束：只支持系统动态分配，不支持指定
	MacAddress string `json:"mac_address"`

	// 功能说明：端口的名称 取值范围：0-255个字符
	Name string `json:"name"`

	// 端口所属网络ID
	NetworkId string `json:"network_id"`

	// 功能说明：端口安全使能标记，如果不使能则安全组和dhcp防欺骗不生效 取值范围：启用（true）或禁用（false）
	PortSecurityEnabled bool `json:"port_security_enabled"`

	// 功能说明：作用在该端口上的安全组的ID列表 约束：不支持更新为空
	SecurityGroups []string `json:"security_groups"`

	// 功能说明：端口状态 取值范围：ACTIVE，BUILD，DOWN 约束：Hana硬直通虚拟机端口状态总为DOWN
	Status string `json:"status"`

	// 项目ID
	TenantId string `json:"tenant_id"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 功能说明：资源创建UTC时间 格式：yyyy-MM-ddTHH:mm:ss
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 功能说明：资源更新UTC时间 格式：yyyy-MM-ddTHH:mm:ss
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o NeutronPort) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronPort struct{}"
	}

	return strings.Join([]string{"NeutronPort", string(data)}, " ")
}
