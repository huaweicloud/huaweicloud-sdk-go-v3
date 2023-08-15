package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FirewallRule 网络ACL规则。
type FirewallRule struct {

	// 策略是否允许  取值范围：allow，deny，reject
	Action FirewallRuleAction `json:"action"`

	// 网络ACL规则描述。
	Description *string `json:"description,omitempty"`

	// 目的地IP地址，IPv4[或IPv6](tag:hide)的CIDR格式
	DestinationIpAddress string `json:"destination_ip_address"`

	// 目的地端口范围  取值范围：整数，比如80，或者以\"-\"隔开的范围，比如80-90
	DestinationPort string `json:"destination_port"`

	// 网络ACL规则使能开关。  取值范围：true，false
	Enabled bool `json:"enabled"`

	// 网络ACL规则ID。  进行更新规则时，如果operate_type为add，则该值为空。
	Id string `json:"id"`

	// IP协议版本  取值范围：4[、6](tag:hide)
	IpVersion int32 `json:"ip_version"`

	// 网络ACL规则名称。
	Name string `json:"name"`

	// 网络ACL规则操作状态，作为请求时取值为\"add\"/\"modify\"/\"delete\"，作为返回值时为\"normal\"。 当请求更新规则时，本参数值为delete时，除id之外，本请求体其他参数均可为空。
	OperateType FirewallRuleOperateType `json:"operate_type"`

	// IP协议，为any时代表所有协议  取值范围：icmp，tcp，udp，[icmpv6，](tag:hide)any
	Protocol FirewallRuleProtocol `json:"protocol"`

	// 源IP地址，IPv4[或IPv6](tag:hide)的CIDR格式
	SourceIpAddress string `json:"source_ip_address"`

	// 源地端口范围  取值范围：整数，比如80，或者以\"-\"隔开的范围，比如80-90
	SourcePort string `json:"source_port"`
}

func (o FirewallRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FirewallRule struct{}"
	}

	return strings.Join([]string{"FirewallRule", string(data)}, " ")
}

type FirewallRuleAction struct {
	value string
}

type FirewallRuleActionEnum struct {
	ALLOW  FirewallRuleAction
	DENY   FirewallRuleAction
	REJECT FirewallRuleAction
}

func GetFirewallRuleActionEnum() FirewallRuleActionEnum {
	return FirewallRuleActionEnum{
		ALLOW: FirewallRuleAction{
			value: "allow",
		},
		DENY: FirewallRuleAction{
			value: "deny",
		},
		REJECT: FirewallRuleAction{
			value: "reject",
		},
	}
}

func (c FirewallRuleAction) Value() string {
	return c.value
}

func (c FirewallRuleAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FirewallRuleAction) UnmarshalJSON(b []byte) error {
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

type FirewallRuleOperateType struct {
	value string
}

type FirewallRuleOperateTypeEnum struct {
	ADD    FirewallRuleOperateType
	MODIFY FirewallRuleOperateType
	DELETE FirewallRuleOperateType
}

func GetFirewallRuleOperateTypeEnum() FirewallRuleOperateTypeEnum {
	return FirewallRuleOperateTypeEnum{
		ADD: FirewallRuleOperateType{
			value: "add",
		},
		MODIFY: FirewallRuleOperateType{
			value: "modify",
		},
		DELETE: FirewallRuleOperateType{
			value: "delete",
		},
	}
}

func (c FirewallRuleOperateType) Value() string {
	return c.value
}

func (c FirewallRuleOperateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FirewallRuleOperateType) UnmarshalJSON(b []byte) error {
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

type FirewallRuleProtocol struct {
	value string
}

type FirewallRuleProtocolEnum struct {
	ICMP FirewallRuleProtocol
	TCP  FirewallRuleProtocol
	UDP  FirewallRuleProtocol
	ANY  FirewallRuleProtocol
}

func GetFirewallRuleProtocolEnum() FirewallRuleProtocolEnum {
	return FirewallRuleProtocolEnum{
		ICMP: FirewallRuleProtocol{
			value: "icmp",
		},
		TCP: FirewallRuleProtocol{
			value: "tcp",
		},
		UDP: FirewallRuleProtocol{
			value: "udp",
		},
		ANY: FirewallRuleProtocol{
			value: "any ",
		},
	}
}

func (c FirewallRuleProtocol) Value() string {
	return c.value
}

func (c FirewallRuleProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FirewallRuleProtocol) UnmarshalJSON(b []byte) error {
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
