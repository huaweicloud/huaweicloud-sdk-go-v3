package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddRuleAclDtoRules **参数解释**： 规则体 **约束限制**： 不涉及
type AddRuleAclDtoRules struct {

	// **参数解释**： 规则名称，由用户定义，用于标识规则 **约束限制**： 字符串长度为0到255 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name string `json:"name"`

	Sequence *OrderRuleAclDto `json:"sequence"`

	// **参数解释**： IP地址的互联网协议类型，用于指定IP地址的互联网协议，由客户指定 **约束限制**： 不涉及 **取值范围**： 0表示IPv4，1表示IPv6 **默认取值**： 不涉及
	AddressType AddRuleAclDtoRulesAddressType `json:"address_type"`

	// **参数解释**： 规则动作类型，用于区分规则对流量的动作 **约束限制**： 仅能使用数字0和1 **取值范围**： 0表示允许通行（permit），1表示拒绝通行（deny） **默认取值**： 不涉及
	ActionType int32 `json:"action_type"`

	// **参数解释**： 规则启用状态，用于区分规则是否启用 **约束限制**： 仅能使用数字0和1 **取值范围**： 0表示启用，1表示禁用 **默认取值**： 不涉及
	Status AddRuleAclDtoRulesStatus `json:"status"`

	// **参数解释**： 规则应用协议列表 **约束限制**： 不涉及 **取值范围**： 规则应用类型包括：“HTTP”，\"HTTPS\"，\"TLS1\"，“DNS”，“SSH”，“MYSQL”，“SMTP”，“RDP”，“RDPS”，“VNC”，“POP3”，“IMAP4”，“SMTPS”，“POP3S”，“FTPS”，“ANY”，“BGP”等。 **默认取值**： 不涉及
	Applications *[]string `json:"applications,omitempty"`

	// **参数解释**： 长连接时长（s），用于表示流量产生会话保持的最大时长。 **约束限制**： 仅能为数字 **取值范围**： 1-86400000。 **默认取值**： 不涉及
	LongConnectTime *int64 `json:"long_connect_time,omitempty"`

	// **参数解释**： 长连接时长对应小时数（h）。 **约束限制**： 仅能为数字 **取值范围**： 0-24000。 **默认取值**： 不涉及
	LongConnectTimeHour *int64 `json:"long_connect_time_hour,omitempty"`

	// **参数解释**： 长连接时长对应分钟数（min）。 **约束限制**： 仅能为数字 **取值范围**： 0-60。 **默认取值**： 不涉及
	LongConnectTimeMinute *int64 `json:"long_connect_time_minute,omitempty"`

	// **参数解释**： 长连接时长对应秒数（s）。 **约束限制**： 仅能为数字 **取值范围**： 0-60。 **默认取值**： 不涉及
	LongConnectTimeSecond *int64 `json:"long_connect_time_second,omitempty"`

	// **参数解释**： 用于表示是否支持长连接。 **约束限制**： 不涉及 **取值范围**： 0表示不支持，1表示支持 **默认取值**： 不涉及
	LongConnectEnable AddRuleAclDtoRulesLongConnectEnable `json:"long_connect_enable"`

	// **参数解释**： 规则描述，用于描述规则的用途。 **约束限制**： 不涉及 **取值范围**： 长度在0-255之间 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 规则方向，用于指定规则是从云上至云下，还是云下至云上 **约束限制**： 当规则type=0（互联网规则）或者type= 2（NAT规则）时方向值必填 **取值范围**： 0表示外到内（云下到云上），1表示内到外（云上到云下）， **默认取值**： 不涉及
	Direction *AddRuleAclDtoRulesDirection `json:"direction,omitempty"`

	Source *RuleAddressDtoForRequest `json:"source"`

	Destination *RuleAddressDtoForRequest `json:"destination"`

	Service *RuleServiceDto `json:"service"`

	Tag *TagsVo `json:"tag,omitempty"`
}

func (o AddRuleAclDtoRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRuleAclDtoRules struct{}"
	}

	return strings.Join([]string{"AddRuleAclDtoRules", string(data)}, " ")
}

type AddRuleAclDtoRulesAddressType struct {
	value int32
}

type AddRuleAclDtoRulesAddressTypeEnum struct {
	E_0 AddRuleAclDtoRulesAddressType
	E_1 AddRuleAclDtoRulesAddressType
	E_2 AddRuleAclDtoRulesAddressType
}

func GetAddRuleAclDtoRulesAddressTypeEnum() AddRuleAclDtoRulesAddressTypeEnum {
	return AddRuleAclDtoRulesAddressTypeEnum{
		E_0: AddRuleAclDtoRulesAddressType{
			value: 0,
		}, E_1: AddRuleAclDtoRulesAddressType{
			value: 1,
		}, E_2: AddRuleAclDtoRulesAddressType{
			value: 2,
		},
	}
}

func (c AddRuleAclDtoRulesAddressType) Value() int32 {
	return c.value
}

func (c AddRuleAclDtoRulesAddressType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddRuleAclDtoRulesAddressType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type AddRuleAclDtoRulesStatus struct {
	value int32
}

type AddRuleAclDtoRulesStatusEnum struct {
	E_0 AddRuleAclDtoRulesStatus
	E_1 AddRuleAclDtoRulesStatus
}

func GetAddRuleAclDtoRulesStatusEnum() AddRuleAclDtoRulesStatusEnum {
	return AddRuleAclDtoRulesStatusEnum{
		E_0: AddRuleAclDtoRulesStatus{
			value: 0,
		}, E_1: AddRuleAclDtoRulesStatus{
			value: 1,
		},
	}
}

func (c AddRuleAclDtoRulesStatus) Value() int32 {
	return c.value
}

func (c AddRuleAclDtoRulesStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddRuleAclDtoRulesStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type AddRuleAclDtoRulesLongConnectEnable struct {
	value int32
}

type AddRuleAclDtoRulesLongConnectEnableEnum struct {
	E_0 AddRuleAclDtoRulesLongConnectEnable
	E_1 AddRuleAclDtoRulesLongConnectEnable
}

func GetAddRuleAclDtoRulesLongConnectEnableEnum() AddRuleAclDtoRulesLongConnectEnableEnum {
	return AddRuleAclDtoRulesLongConnectEnableEnum{
		E_0: AddRuleAclDtoRulesLongConnectEnable{
			value: 0,
		}, E_1: AddRuleAclDtoRulesLongConnectEnable{
			value: 1,
		},
	}
}

func (c AddRuleAclDtoRulesLongConnectEnable) Value() int32 {
	return c.value
}

func (c AddRuleAclDtoRulesLongConnectEnable) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddRuleAclDtoRulesLongConnectEnable) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type AddRuleAclDtoRulesDirection struct {
	value int32
}

type AddRuleAclDtoRulesDirectionEnum struct {
	E_0 AddRuleAclDtoRulesDirection
	E_1 AddRuleAclDtoRulesDirection
}

func GetAddRuleAclDtoRulesDirectionEnum() AddRuleAclDtoRulesDirectionEnum {
	return AddRuleAclDtoRulesDirectionEnum{
		E_0: AddRuleAclDtoRulesDirection{
			value: 0,
		}, E_1: AddRuleAclDtoRulesDirection{
			value: 1,
		},
	}
}

func (c AddRuleAclDtoRulesDirection) Value() int32 {
	return c.value
}

func (c AddRuleAclDtoRulesDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddRuleAclDtoRulesDirection) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
