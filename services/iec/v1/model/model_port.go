package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Port 端口的字典对象。
type Port struct {

	// 端口唯一标识
	Id *string `json:"id,omitempty"`

	// 端口名称  取值：默认为空，最大长度不超过255
	Name *string `json:"name,omitempty"`

	// 端口状态，Hana硬直通虚拟机端口状态总为DOWN  取值范围：ACTIVE、BUILD、DOWN
	Status *PortStatus `json:"status,omitempty"`

	// 管理状态  约束：只支持true，默认为true
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 端口IP。  约束：一个端口只支持一个fixed_ip，且不支持更新。
	FixedIps *[]FixedIp `json:"fixed_ips,omitempty"`

	// 端口MAC地址  约束：由系统分配，不支持指定
	MacAddress *string `json:"mac_address,omitempty"`

	// 端口所属网络的ID  约束：必须是存在的网络ID
	NetworkId *string `json:"network_id,omitempty"`

	// 端口所属设备ID  约束：不支持设置和更新，由系统自动维护
	DeviceId *string `json:"device_id,omitempty"`

	// 设备所属（DHCP/Router/ lb/Nova）  约束：不支持设置和更新，由系统自动维护
	DeviceOwner *string `json:"device_owner,omitempty"`

	// 安全组的UUID(扩展属性)
	SecurityGroups *[]string `json:"security_groups,omitempty"`

	// DHCP的扩展属性。
	ExtraDhcpOpts *[]ExtraDhcpOption `json:"extra_dhcp_opts,omitempty"`

	// IP/Mac对列表。  约束：IP地址不允许为 “0.0.0.0/0”  建议：如果allowed_address_pairs配置地址池较大的CIDR（掩码小于24位），建议为该port配置一个单独的安全组。
	AllowedAddressPairs *[]AllowedAddressPair `json:"allowed_address_pairs,omitempty"`

	// 站点ID
	SiteId *string `json:"site_id,omitempty"`

	// 主网卡默认内网域名信息  约束：不支持设置和更新，由系统自动维护
	DnsAssignment *[]DnsAssignment `json:"dns_assignment,omitempty"`

	// 主网卡默认内网DNS名称  约束：不支持设置和更新，由系统自动维护
	DnsName *string `json:"dns_name,omitempty"`
}

func (o Port) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Port struct{}"
	}

	return strings.Join([]string{"Port", string(data)}, " ")
}

type PortStatus struct {
	value string
}

type PortStatusEnum struct {
	ACTIVE PortStatus
	BUILD  PortStatus
	DOWN   PortStatus
}

func GetPortStatusEnum() PortStatusEnum {
	return PortStatusEnum{
		ACTIVE: PortStatus{
			value: "ACTIVE",
		},
		BUILD: PortStatus{
			value: "BUILD",
		},
		DOWN: PortStatus{
			value: "DOWN",
		},
	}
}

func (c PortStatus) Value() string {
	return c.value
}

func (c PortStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PortStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
