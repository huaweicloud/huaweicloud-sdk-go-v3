package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateRuleAclDto UpdateRuleAclDto
type UpdateRuleAclDto struct {

	// **参数解释**： IP地址的互联网协议类型，用于指定IP地址的互联网协议，由客户指定 **约束限制**： 不涉及 **取值范围**： 0表示IPv4，1表示IPv6 **默认取值**： 不涉及
	AddressType *UpdateRuleAclDtoAddressType `json:"address_type,omitempty"`

	// **参数解释**： 规则名称，由用户定义，用于标识规则 **约束限制**： 字符串长度为0到255 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 规则方向，用于指定规则是从云上至云下，还是云下至云上 **约束限制**： 当规则type=0（互联网规则）或者type= 2（nat规则）时方向值必填 **取值范围**： 0表示外到内（云下到云上），1表示内到外（云上到云下）， **默认取值**： 不涉及
	Direction *UpdateRuleAclDtoDirection `json:"direction,omitempty"`

	// **参数解释**： 规则动作类型，用于区分规则对流量的动作 **约束限制**： 仅能使用数字0和1 **取值范围**： 0表示允许通行（permit），1表示拒绝通行（deny） **默认取值**： 不涉及
	ActionType *UpdateRuleAclDtoActionType `json:"action_type,omitempty"`

	// **参数解释**： 规则启用状态，用于区分规则是否启用 **约束限制**： 仅能使用数字0和1 **取值范围**： 0表示启用，1表示禁用 **默认取值**： 不涉及
	Status *int32 `json:"status,omitempty"`

	// **参数解释**： 规则应用协议列表 **约束限制**： 不涉及 **取值范围**： 规则应用类型包括：“HTTP”，\"HTTPS\"，\"TLS1\"，“DNS”，“SSH”，“MYSQL”，“SMTP”，“RDP”，“RDPS”，“VNC”，“POP3”，“IMAP4”，“SMTPS”，“POP3S”，“FTPS”，“ANY”，“BGP”等。 **默认取值**： 不涉及
	Applications *[]string `json:"applications,omitempty"`

	// **参数解释**： 规则描述，用于描述规则的用途。 **约束限制**： 不涉及 **取值范围**： 长度在0-255之间 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 长连接时长对应小时数（h）。 **约束限制**： 仅能为数字 **取值范围**： 0-24000。 **默认取值**： 不涉及
	LongConnectTimeHour *int64 `json:"long_connect_time_hour,omitempty"`

	// **参数解释**： 长连接时长对应分钟数（min）。 **约束限制**： 仅能为数字 **取值范围**： 0-60。 **默认取值**： 不涉及
	LongConnectTimeMinute *int64 `json:"long_connect_time_minute,omitempty"`

	// **参数解释**： 长连接时长对应秒数（s）。 **约束限制**： 仅能为数字 **取值范围**： 0-60。 **默认取值**： 不涉及
	LongConnectTimeSecond *int64 `json:"long_connect_time_second,omitempty"`

	// **参数解释**： 长连接时长（s），用于表示流量产生会话保持的最大时长。 **约束限制**： 仅能为数字 **取值范围**： 1-86400000。 **默认取值**： 不涉及
	LongConnectTime *int64 `json:"long_connect_time,omitempty"`

	// **参数解释**： 用于表示是否支持长连接。 **约束限制**： 不涉及 **取值范围**： 0表示不支持，1表示支持 **默认取值**： 不涉及
	LongConnectEnable *UpdateRuleAclDtoLongConnectEnable `json:"long_connect_enable,omitempty"`

	Source *RuleAddressDto `json:"source,omitempty"`

	Destination *RuleAddressDto `json:"destination,omitempty"`

	Service *RuleServiceDto `json:"service,omitempty"`

	// **参数解释**： 规则类型，用于区分不同防护对象设置规则类型。 **约束限制**： 不涉及 **取值范围**：  0：互联网边界规则，源（source）和目的（destination）地址需要为公网IP或域名； 1：VPC间规则，源（source）和目的（destination）地址需要为私有ip； 2：NAT规则，源（source）地址需要为私网IP，目的地址为公网IP或域名。 **默认取值**： 不涉及
	Type *UpdateRuleAclDtoType `json:"type,omitempty"`

	Tag *TagsVo `json:"tag,omitempty"`
}

func (o UpdateRuleAclDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleAclDto struct{}"
	}

	return strings.Join([]string{"UpdateRuleAclDto", string(data)}, " ")
}

type UpdateRuleAclDtoAddressType struct {
	value int32
}

type UpdateRuleAclDtoAddressTypeEnum struct {
	E_0 UpdateRuleAclDtoAddressType
	E_1 UpdateRuleAclDtoAddressType
}

func GetUpdateRuleAclDtoAddressTypeEnum() UpdateRuleAclDtoAddressTypeEnum {
	return UpdateRuleAclDtoAddressTypeEnum{
		E_0: UpdateRuleAclDtoAddressType{
			value: 0,
		}, E_1: UpdateRuleAclDtoAddressType{
			value: 1,
		},
	}
}

func (c UpdateRuleAclDtoAddressType) Value() int32 {
	return c.value
}

func (c UpdateRuleAclDtoAddressType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRuleAclDtoAddressType) UnmarshalJSON(b []byte) error {
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

type UpdateRuleAclDtoDirection struct {
	value int32
}

type UpdateRuleAclDtoDirectionEnum struct {
	E_0 UpdateRuleAclDtoDirection
	E_1 UpdateRuleAclDtoDirection
}

func GetUpdateRuleAclDtoDirectionEnum() UpdateRuleAclDtoDirectionEnum {
	return UpdateRuleAclDtoDirectionEnum{
		E_0: UpdateRuleAclDtoDirection{
			value: 0,
		}, E_1: UpdateRuleAclDtoDirection{
			value: 1,
		},
	}
}

func (c UpdateRuleAclDtoDirection) Value() int32 {
	return c.value
}

func (c UpdateRuleAclDtoDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRuleAclDtoDirection) UnmarshalJSON(b []byte) error {
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

type UpdateRuleAclDtoActionType struct {
	value int32
}

type UpdateRuleAclDtoActionTypeEnum struct {
	E_0 UpdateRuleAclDtoActionType
	E_1 UpdateRuleAclDtoActionType
}

func GetUpdateRuleAclDtoActionTypeEnum() UpdateRuleAclDtoActionTypeEnum {
	return UpdateRuleAclDtoActionTypeEnum{
		E_0: UpdateRuleAclDtoActionType{
			value: 0,
		}, E_1: UpdateRuleAclDtoActionType{
			value: 1,
		},
	}
}

func (c UpdateRuleAclDtoActionType) Value() int32 {
	return c.value
}

func (c UpdateRuleAclDtoActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRuleAclDtoActionType) UnmarshalJSON(b []byte) error {
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

type UpdateRuleAclDtoLongConnectEnable struct {
	value int32
}

type UpdateRuleAclDtoLongConnectEnableEnum struct {
	E_0 UpdateRuleAclDtoLongConnectEnable
	E_1 UpdateRuleAclDtoLongConnectEnable
}

func GetUpdateRuleAclDtoLongConnectEnableEnum() UpdateRuleAclDtoLongConnectEnableEnum {
	return UpdateRuleAclDtoLongConnectEnableEnum{
		E_0: UpdateRuleAclDtoLongConnectEnable{
			value: 0,
		}, E_1: UpdateRuleAclDtoLongConnectEnable{
			value: 1,
		},
	}
}

func (c UpdateRuleAclDtoLongConnectEnable) Value() int32 {
	return c.value
}

func (c UpdateRuleAclDtoLongConnectEnable) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRuleAclDtoLongConnectEnable) UnmarshalJSON(b []byte) error {
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

type UpdateRuleAclDtoType struct {
	value int32
}

type UpdateRuleAclDtoTypeEnum struct {
	E_0 UpdateRuleAclDtoType
	E_1 UpdateRuleAclDtoType
	E_2 UpdateRuleAclDtoType
}

func GetUpdateRuleAclDtoTypeEnum() UpdateRuleAclDtoTypeEnum {
	return UpdateRuleAclDtoTypeEnum{
		E_0: UpdateRuleAclDtoType{
			value: 0,
		}, E_1: UpdateRuleAclDtoType{
			value: 1,
		}, E_2: UpdateRuleAclDtoType{
			value: 2,
		},
	}
}

func (c UpdateRuleAclDtoType) Value() int32 {
	return c.value
}

func (c UpdateRuleAclDtoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRuleAclDtoType) UnmarshalJSON(b []byte) error {
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
