package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RuleAclListResponseDtoDataRecords items
type RuleAclListResponseDtoDataRecords struct {

	// **参数解释**： 规则ID **取值范围**： 不涉及
	RuleId *string `json:"rule_id,omitempty"`

	// **参数解释**： 排序id **取值范围**： 不涉及
	OrderId *int32 `json:"order_id,omitempty"`

	// **参数解释**： 应用列表 **取值范围**： 不涉及
	Applications *[]string `json:"applications,omitempty"`

	// 参数解释： IP地址的互联网协议类型，用于指定IP地址的互联网协议，由客户指定 约束限制： 不涉及 取值范围： 0：IPv4 1：IPv6 默认取值： 不涉及
	AddressType *int32 `json:"address_type,omitempty"`

	// **参数解释**： 规则名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 规则方向 **取值范围**： 0：外到内1：内到外
	Direction *RuleAclListResponseDtoDataRecordsDirection `json:"direction,omitempty"`

	// **参数解释**： 规则动作类型，用于区分规则对流量的动作 **取值范围**： 0表示允许通行（permit），1表示拒绝通行（deny）
	ActionType *int32 `json:"action_type,omitempty"`

	// **参数解释**： 规则启用状态，用于区分规则是否启用 **取值范围**： 0表示启用，1表示禁用
	Status *int32 `json:"status,omitempty"`

	// **参数解释**： 规则描述，用于描述规则的用途。 **取值范围**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 长连接时长（s）。 **取值范围**： 1-86400000。
	LongConnectTime *int64 `json:"long_connect_time,omitempty"`

	// **参数解释**： 用于表示是否支持长连接。 **取值范围**： 0表示不支持，1表示支持
	LongConnectEnable *int32 `json:"long_connect_enable,omitempty"`

	// **参数解释**： 长连接时长对应小时数（h）。 **取值范围**： 0-24000。
	LongConnectTimeHour *int64 `json:"long_connect_time_hour,omitempty"`

	// **参数解释**： 长连接时长对应分钟数（min）。 **取值范围**： 0-60。
	LongConnectTimeMinute *int64 `json:"long_connect_time_minute,omitempty"`

	// **参数解释**： 长连接时长对应秒数（s）。 **取值范围**： 0-60。
	LongConnectTimeSecond *int64 `json:"long_connect_time_second,omitempty"`

	Source *RuleAddressDtoForResponse `json:"source,omitempty"`

	Destination *RuleAddressDtoForResponse `json:"destination,omitempty"`

	Service *RuleServiceDtoForResponse `json:"service,omitempty"`

	// **参数解释**： 规则类型，用于区分不同防护对象设置规则类型。 **取值范围**：  0：互联网边界规则，源（source）和目的（destination）地址需要为公网IP或域名； 1：VPC间规则，源（source）和目的（destination）地址需要为私有ip； 2：NAT规则，源（source）地址需要为私网IP，目的地址为公网IP或域名。
	Type *RuleAclListResponseDtoDataRecordsType `json:"type,omitempty"`

	// **参数解释**： 规则创建时间。 **取值范围**： 不涉及
	CreatedDate *string `json:"created_date,omitempty"`

	// **参数解释**： 规则修改时间。 **取值范围**： 不涉及
	ModifiedDate *string `json:"modified_date,omitempty"`

	// **参数解释**： 规则最后开启时间。 **取值范围**： 不涉及
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
