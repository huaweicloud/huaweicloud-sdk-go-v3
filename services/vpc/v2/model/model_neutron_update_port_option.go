package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdatePortOption
type NeutronUpdatePortOption struct {

	// 功能说明：网络的名称 取值范围：0-255个字符
	Name *string `json:"name,omitempty"`

	// 功能说明：扩展属性：安全组的UUID 例如：\"security_groups\": [\"a0608cbf-d047-4f54-8b28-cd7b59853fff\"] 约束：不支持更新为空
	SecurityGroups *[]string `json:"security_groups,omitempty"`

	// 功能说明：扩展属性：IP/Mac对列表，allow_address_pair参见“allow_address_pair对象”表 约束：  - IP地址不允许为 “0.0.0.0”  - 如果allowed_address_pairs配置地址池较大的CIDR（掩码小于24位），建议为该port配置一个单独的安全组  - 硬件SDN环境不支持ip_address属性配置为CIDR格式  - 为虚拟IP配置后端ECS场景，allowed_address_pairs中配置的IP地址，必须为ECS网卡已有的IP地址，否则可能会导致虚拟IP通信异常。
	AllowedAddressPairs *[]AllowedAddressPair `json:"allowed_address_pairs,omitempty"`

	// 扩展属性：DHCP的扩展Option
	ExtraDhcpOpts *[]ExtraDhcpOpt `json:"extra_dhcp_opts,omitempty"`

	// 功能说明：端口安全使能标记，如果不使能则安全组和dhcp防欺骗不生效，默认为true
	PortSecurityEnabled *bool `json:"port_security_enabled,omitempty"`

	// 绑定的vNIC类型  - normal: 软交换
	BindingvnicType *string `json:"binding:vnic_type,omitempty"`

	// 功能说明：扩展属性，提供用户设置自定义信息  - internal_elb字段，布尔类型，普通租户可见。只有在创建内网ELB的虚拟IP的网卡时设置为true。普通租户没有权限更改该字段，由系统维护。 举例：{\"internal_elb\": true}  - disable_security_groups字段，布尔类型，普通租户可见。默认为false，高性能通信场景下，允许指定为true。仅支持创建port和读取时指定。当前仅支持指定为true，不支持指定为false。 举例：{\"disable_security_groups\"：true }，当前仅支持指定为true，不支持指定为false，指定为true时，FWaaS功能不生效。
	Bindingprofile map[string]interface{} `json:"binding:profile,omitempty"`
}

func (o NeutronUpdatePortOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdatePortOption struct{}"
	}

	return strings.Join([]string{"NeutronUpdatePortOption", string(data)}, " ")
}
