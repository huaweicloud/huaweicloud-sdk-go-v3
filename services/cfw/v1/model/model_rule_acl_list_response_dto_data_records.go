package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RuleAclListResponseDtoDataRecords items
type RuleAclListResponseDtoDataRecords struct {

	// 规则id
	RuleId *string `json:"rule_id,omitempty"`

	// 地址类型0 ipv4,1 ipv6
	AddressType *int32 `json:"address_type,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 规则方向0：外到内1：内到外
	Direction *RuleAclListResponseDtoDataRecordsDirection `json:"direction,omitempty"`

	// 动作0：permit,1：deny
	ActionType *int32 `json:"action_type,omitempty"`

	// 规则下发状态 0：禁用,1：启用
	Status *int32 `json:"status,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 长连接时长
	LongConnectTime *int64 `json:"long_connect_time,omitempty"`

	// 长连接支持
	LongConnectEnable *int32 `json:"long_connect_enable,omitempty"`

	// 长连接时长小时
	LongConnectTimeHour *int64 `json:"long_connect_time_hour,omitempty"`

	// 长连接时长分钟
	LongConnectTimeMinute *int64 `json:"long_connect_time_minute,omitempty"`

	// 长连接时长秒
	LongConnectTimeSecond *int64 `json:"long_connect_time_second,omitempty"`

	Source *RuleAddressDtoForResponse `json:"source,omitempty"`

	Destination *RuleAddressDtoForResponse `json:"destination,omitempty"`

	Service *RuleServiceDtoForResponse `json:"service,omitempty"`

	// 规则type，0：互联网规则，1：vpc规则，2：nat规则
	Type *RuleAclListResponseDtoDataRecordsType `json:"type,omitempty"`

	// 创建时间
	CreatedDate *string `json:"created_date,omitempty"`

	// 最后开启时间
	LastOpenTime *string `json:"last_open_time,omitempty"`

	Tag *TagsVo `json:"tag,omitempty"`
}

func (o RuleAclListResponseDtoDataRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleAclListResponseDtoDataRecords struct{}"
	}

	return strings.Join([]string{"RuleAclListResponseDtoDataRecords", string(data)}, " ")
}

type RuleAclListResponseDtoDataRecordsDirection struct {
	value int32
}

type RuleAclListResponseDtoDataRecordsDirectionEnum struct {
	E_0 RuleAclListResponseDtoDataRecordsDirection
	E_1 RuleAclListResponseDtoDataRecordsDirection
}

func GetRuleAclListResponseDtoDataRecordsDirectionEnum() RuleAclListResponseDtoDataRecordsDirectionEnum {
	return RuleAclListResponseDtoDataRecordsDirectionEnum{
		E_0: RuleAclListResponseDtoDataRecordsDirection{
			value: 0,
		}, E_1: RuleAclListResponseDtoDataRecordsDirection{
			value: 1,
		},
	}
}

func (c RuleAclListResponseDtoDataRecordsDirection) Value() int32 {
	return c.value
}

func (c RuleAclListResponseDtoDataRecordsDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuleAclListResponseDtoDataRecordsDirection) UnmarshalJSON(b []byte) error {
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

type RuleAclListResponseDtoDataRecordsType struct {
	value int32
}

type RuleAclListResponseDtoDataRecordsTypeEnum struct {
	E_0 RuleAclListResponseDtoDataRecordsType
	E_1 RuleAclListResponseDtoDataRecordsType
	E_2 RuleAclListResponseDtoDataRecordsType
}

func GetRuleAclListResponseDtoDataRecordsTypeEnum() RuleAclListResponseDtoDataRecordsTypeEnum {
	return RuleAclListResponseDtoDataRecordsTypeEnum{
		E_0: RuleAclListResponseDtoDataRecordsType{
			value: 0,
		}, E_1: RuleAclListResponseDtoDataRecordsType{
			value: 1,
		}, E_2: RuleAclListResponseDtoDataRecordsType{
			value: 2,
		},
	}
}

func (c RuleAclListResponseDtoDataRecordsType) Value() int32 {
	return c.value
}

func (c RuleAclListResponseDtoDataRecordsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuleAclListResponseDtoDataRecordsType) UnmarshalJSON(b []byte) error {
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
