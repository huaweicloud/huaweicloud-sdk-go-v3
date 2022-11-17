package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateRuleAclDto
type UpdateRuleAclDto struct {

	// 地址类型，0 ipv4,1 ipv6
	AddressType *UpdateRuleAclDtoAddressType `json:"address_type,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	Sequence *OrderRuleAclDto `json:"sequence,omitempty"`

	// 规则方向
	Direction *UpdateRuleAclDtoDirection `json:"direction,omitempty"`

	// 动作0：permit,1：deny
	ActionType *UpdateRuleAclDtoActionType `json:"action_type,omitempty"`

	// 规则下发状态 0：禁用,1：启用
	Status *int32 `json:"status,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 长连接时长小时
	LongConnectTimeHour *int64 `json:"long_connect_time_hour,omitempty"`

	// 长连接时长分钟
	LongConnectTimeMinute *int64 `json:"long_connect_time_minute,omitempty"`

	// 长连接时长秒
	LongConnectTimeSecond *int64 `json:"long_connect_time_second,omitempty"`

	// 长连接时长
	LongConnectTime *int64 `json:"long_connect_time,omitempty"`

	// 是否支持长连接，0表示支持，1表示不支持
	LongConnectEnable *UpdateRuleAclDtoLongConnectEnable `json:"long_connect_enable,omitempty"`

	Source *RuleAddressDto `json:"source,omitempty"`

	Destination *RuleAddressDto `json:"destination,omitempty"`

	Service *RuleServiceDto `json:"service,omitempty"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
