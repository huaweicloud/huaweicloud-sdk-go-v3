package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type NeutronUpdateFirewallRuleOption struct {

	// 对通过网络ACL防火墙的流量执行的操作。
	Action *NeutronUpdateFirewallRuleOptionAction `json:"action,omitempty" xml:"action"`

	// 网络ACL防火墙规则描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 目的地址或者CIDR。
	DestinationIpAddress *string `json:"destination_ip_address,omitempty" xml:"destination_ip_address"`

	// 目的端口号或者一段端口范围。
	DestinationPort *string `json:"destination_port,omitempty" xml:"destination_port"`

	// 是否使能网络ACL防火墙规则。
	Enabled *bool `json:"enabled,omitempty" xml:"enabled"`

	// IP协议版本。
	IpVersion *int32 `json:"ip_version,omitempty" xml:"ip_version"`

	// 网络ACL防火墙规则名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// IP协议，支持TCP,UDP,ICMP, ICMPV6或者IP协议号（0-255）
	Protocol *string `json:"protocol,omitempty" xml:"protocol"`

	// 源地址或者CIDR。
	SourceIpAddress *string `json:"source_ip_address,omitempty" xml:"source_ip_address"`

	// 源端口号或者一段端口范围。
	SourcePort *string `json:"source_port,omitempty" xml:"source_port"`
}

func (o NeutronUpdateFirewallRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFirewallRuleOption struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFirewallRuleOption", string(data)}, " ")
}

type NeutronUpdateFirewallRuleOptionAction struct {
	value string
}

type NeutronUpdateFirewallRuleOptionActionEnum struct {
	DENY  NeutronUpdateFirewallRuleOptionAction
	ALLOW NeutronUpdateFirewallRuleOptionAction
}

func GetNeutronUpdateFirewallRuleOptionActionEnum() NeutronUpdateFirewallRuleOptionActionEnum {
	return NeutronUpdateFirewallRuleOptionActionEnum{
		DENY: NeutronUpdateFirewallRuleOptionAction{
			value: "DENY",
		},
		ALLOW: NeutronUpdateFirewallRuleOptionAction{
			value: "ALLOW",
		},
	}
}

func (c NeutronUpdateFirewallRuleOptionAction) Value() string {
	return c.value
}

func (c NeutronUpdateFirewallRuleOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NeutronUpdateFirewallRuleOptionAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
