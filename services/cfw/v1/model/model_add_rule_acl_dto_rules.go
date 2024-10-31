package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddRuleAclDtoRules items
type AddRuleAclDtoRules struct {

	// 规则名称
	Name string `json:"name"`

	Sequence *OrderRuleAclDto `json:"sequence"`

	// 地址类型，0表示ipv4，1表示ipv6
	AddressType AddRuleAclDtoRulesAddressType `json:"address_type"`

	// 规则动作，0表示允许通行（permit），1表示拒绝通行（deny）
	ActionType int32 `json:"action_type"`

	// 规则启用状态，0表示禁用，1表示启用
	Status AddRuleAclDtoRulesStatus `json:"status"`

	// 规则应用列表，规则应用类型包括：“HTTP”，\"HTTPS\"，\"TLS1\"，“DNS”，“SSH”，“MYSQL”，“SMTP”，“RDP”，“RDPS”，“VNC”，“POP3”，“IMAP4”，“SMTPS”，“POP3S”，“FTPS”，“ANY”，“BGP”等。
	Applications *[]string `json:"applications,omitempty"`

	// 长连接时长
	LongConnectTime *int64 `json:"long_connect_time,omitempty"`

	// 长连接时长对应小时
	LongConnectTimeHour *int64 `json:"long_connect_time_hour,omitempty"`

	// 长连接时长对应分钟
	LongConnectTimeMinute *int64 `json:"long_connect_time_minute,omitempty"`

	// 长连接时长秒
	LongConnectTimeSecond *int64 `json:"long_connect_time_second,omitempty"`

	// 是否支持长连接，0表示不支持长连接，1表示支持长连接
	LongConnectEnable AddRuleAclDtoRulesLongConnectEnable `json:"long_connect_enable"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 方向：0表示外到内，1表示内到外，规则type=0（互联网规则）或者type= 2（nat规则）时方向值必填
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
