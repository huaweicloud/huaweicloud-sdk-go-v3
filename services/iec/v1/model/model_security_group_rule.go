package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 安全组规则。
type SecurityGroupRule struct {
	// 安全组规则的ID。

	Id *string `json:"id,omitempty"`
	// 安全组规则描述信息。

	Description *string `json:"description,omitempty"`
	// 安全组ID。

	SecurityGroupId *string `json:"security_group_id,omitempty"`
	// 出入控制方向。  取值范围：  - egress：出方向  - ingress：入方向

	Direction *SecurityGroupRuleDirection `json:"direction,omitempty"`
	// IP协议类型。  取值范围：IPv4[，IPv6](tag:hide)  约束：不填默认值为IPv4

	Ethertype *SecurityGroupRuleEthertype `json:"ethertype,omitempty"`
	// 协议类型。  取值范围：icmp、tcp、udp等  约束：为空表示支持所有协议

	Protocol *string `json:"protocol,omitempty"`
	// 起始端口值。  取值范围：1~65535  约束：取值不能大于port_range_max的值，为空表示所有端口

	PortRangeMin *string `json:"port_range_min,omitempty"`
	// 结束端口值。  取值范围：1~65535  约束：取值不能小于port_range_min的值，为空表示所有端口。

	PortRangeMax *string `json:"port_range_max,omitempty"`
	// 对端安全组ID。  约束：和remote_ip_prefix互斥 ，remote_group_id与remote_ip_prefix必须存在一个

	RemoteGroupId *string `json:"remote_group_id,omitempty"`
	// 远端IP地址，当direction是egress时为虚拟机访问端的地址，当direction是ingress时为访问虚拟机的地址。  取值范围：IP地址，或者cidr格式  约束：和remote_group_id互斥

	RemoteIpPrefix *string `json:"remote_ip_prefix,omitempty"`
}

func (o SecurityGroupRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupRule struct{}"
	}

	return strings.Join([]string{"SecurityGroupRule", string(data)}, " ")
}

type SecurityGroupRuleDirection struct {
	value string
}

type SecurityGroupRuleDirectionEnum struct {
	INGRESS SecurityGroupRuleDirection
	EGRESS  SecurityGroupRuleDirection
}

func GetSecurityGroupRuleDirectionEnum() SecurityGroupRuleDirectionEnum {
	return SecurityGroupRuleDirectionEnum{
		INGRESS: SecurityGroupRuleDirection{
			value: "ingress",
		},
		EGRESS: SecurityGroupRuleDirection{
			value: "egress",
		},
	}
}

func (c SecurityGroupRuleDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SecurityGroupRuleDirection) UnmarshalJSON(b []byte) error {
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

type SecurityGroupRuleEthertype struct {
	value string
}

type SecurityGroupRuleEthertypeEnum struct {
	I_PV4 SecurityGroupRuleEthertype
}

func GetSecurityGroupRuleEthertypeEnum() SecurityGroupRuleEthertypeEnum {
	return SecurityGroupRuleEthertypeEnum{
		I_PV4: SecurityGroupRuleEthertype{
			value: "IPv4",
		},
	}
}

func (c SecurityGroupRuleEthertype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SecurityGroupRuleEthertype) UnmarshalJSON(b []byte) error {
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
