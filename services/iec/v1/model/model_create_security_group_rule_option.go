package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSecurityGroupRuleOption 创建安全组规则参数
type CreateSecurityGroupRuleOption struct {

	// 安全组规则描述信息。
	Description *string `json:"description,omitempty"`

	// 安全组ID。
	SecurityGroupId string `json:"security_group_id"`

	// 出入控制方向。  取值范围：  - egress：出方向  - ingress：入方向
	Direction CreateSecurityGroupRuleOptionDirection `json:"direction"`

	// IP协议类型。  取值范围：IPv4,IPv6
	Ethertype *CreateSecurityGroupRuleOptionEthertype `json:"ethertype,omitempty"`

	// 协议类型。  取值范围：icmp、tcp、udp、icmpv6或IP协议号（0~255）  约束：为空表示支持所有协议
	Protocol *string `json:"protocol,omitempty"`

	// 起始端口值。  取值范围：1~65535  约束：取值不能大于port_range_max的值，为空表示所有端口。
	PortRangeMin *int32 `json:"port_range_min,omitempty"`

	// 结束端口值。  取值范围：1~65535  约束：取值不能小于port_range_min的值，为空表示所有端口。
	PortRangeMax *int32 `json:"port_range_max,omitempty"`

	// 对端安全组id。  约束：和remote_ip_prefix互斥
	RemoteGroupId *string `json:"remote_group_id,omitempty"`

	// 远端IP地址，当direction是egress时为虚拟机访问端的地址，当direction是ingress时为访问虚拟机的地址。  取值范围：IP地址，或者cidr格式  约束：和remote_group_id互斥
	RemoteIpPrefix *string `json:"remote_ip_prefix,omitempty"`

	// 安全组规则生效策略 取值范围：allow 允许，deny 拒绝  约束：默认值为allow
	Action *string `json:"action,omitempty"`

	// 规则在安全组中的优先级  取值范围：1~100，1代表最高优先级  约束：默认值为1
	Priority *int32 `json:"priority,omitempty"`
}

func (o CreateSecurityGroupRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupRuleOption struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupRuleOption", string(data)}, " ")
}

type CreateSecurityGroupRuleOptionDirection struct {
	value string
}

type CreateSecurityGroupRuleOptionDirectionEnum struct {
	INGRESS CreateSecurityGroupRuleOptionDirection
	EGRESS  CreateSecurityGroupRuleOptionDirection
}

func GetCreateSecurityGroupRuleOptionDirectionEnum() CreateSecurityGroupRuleOptionDirectionEnum {
	return CreateSecurityGroupRuleOptionDirectionEnum{
		INGRESS: CreateSecurityGroupRuleOptionDirection{
			value: "ingress",
		},
		EGRESS: CreateSecurityGroupRuleOptionDirection{
			value: "egress",
		},
	}
}

func (c CreateSecurityGroupRuleOptionDirection) Value() string {
	return c.value
}

func (c CreateSecurityGroupRuleOptionDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecurityGroupRuleOptionDirection) UnmarshalJSON(b []byte) error {
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

type CreateSecurityGroupRuleOptionEthertype struct {
	value string
}

type CreateSecurityGroupRuleOptionEthertypeEnum struct {
	I_PV4 CreateSecurityGroupRuleOptionEthertype
	I_PV6 CreateSecurityGroupRuleOptionEthertype
}

func GetCreateSecurityGroupRuleOptionEthertypeEnum() CreateSecurityGroupRuleOptionEthertypeEnum {
	return CreateSecurityGroupRuleOptionEthertypeEnum{
		I_PV4: CreateSecurityGroupRuleOptionEthertype{
			value: "IPv4",
		},
		I_PV6: CreateSecurityGroupRuleOptionEthertype{
			value: "IPv6",
		},
	}
}

func (c CreateSecurityGroupRuleOptionEthertype) Value() string {
	return c.value
}

func (c CreateSecurityGroupRuleOptionEthertype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSecurityGroupRuleOptionEthertype) UnmarshalJSON(b []byte) error {
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
