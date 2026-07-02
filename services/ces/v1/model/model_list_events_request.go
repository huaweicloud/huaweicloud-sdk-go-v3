package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEventsRequest Request Object
type ListEventsRequest struct {

	// **参数解释** 事件类型 **约束限制** 不涉及 **取值范围** 枚举值： - EVENT.SYS 系统事件 - EVENT.CUSTOM 自定义事件 **默认取值** 不涉及
	EventType *ListEventsRequestEventType `json:"event_type,omitempty"`

	// **参数解释** 事件子类型 **约束限制** 不涉及 **取值范围** 枚举值： - SUB_EVENT.OPS 运维事件 - SUB_EVENT.PLAN 计划事件 - SUB_EVENT.CUSTOM 自定义事件 **默认取值** 不涉及
	SubEventType *ListEventsRequestSubEventType `json:"sub_event_type,omitempty"`

	// **参数解释** 事件名称，值为系统产生的事件名称，或用户自定义上报的事件名称 **约束限制** 不涉及 **取值范围** 必须以字母（大写或小写）开头，后面可以跟零个或多个字母（大写或小写）、数字、下划线（_），长度为[1,64]个字符 **默认取值** 不涉及
	EventName *string `json:"event_name,omitempty"`

	// **参数解释** 查询数据起始时间，UNIX时间戳，单位毫秒 **约束限制** 不涉及 **取值范围** 毫秒级时间戳，范围为[1111111111111,9999999999999] **默认取值** 不涉及
	From *int64 `json:"from,omitempty"`

	// **参数解释** 查询数据截止时间，UNIX时间戳，单位毫秒 **约束限制** from必须小于to **取值范围** 毫秒级时间戳，范围为[1111111111111,9999999999999] **默认取值** 不涉及
	To *int64 `json:"to,omitempty"`

	// **参数解释** 分页起始值 **约束限制** 不涉及 **取值范围** 非负整数 **默认取值** 0
	Start *string `json:"start,omitempty"`

	// **参数解释** 单次查询的条数限制 **约束限制** 不涉及 **取值范围** 条数限制为[1,100] **默认取值** 100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsRequest struct{}"
	}

	return strings.Join([]string{"ListEventsRequest", string(data)}, " ")
}

type ListEventsRequestEventType struct {
	value string
}

type ListEventsRequestEventTypeEnum struct {
	EVENT_SYS    ListEventsRequestEventType
	EVENT_CUSTOM ListEventsRequestEventType
}

func GetListEventsRequestEventTypeEnum() ListEventsRequestEventTypeEnum {
	return ListEventsRequestEventTypeEnum{
		EVENT_SYS: ListEventsRequestEventType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: ListEventsRequestEventType{
			value: "EVENT.CUSTOM",
		},
	}
}

func (c ListEventsRequestEventType) Value() string {
	return c.value
}

func (c ListEventsRequestEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventsRequestEventType) UnmarshalJSON(b []byte) error {
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

type ListEventsRequestSubEventType struct {
	value string
}

type ListEventsRequestSubEventTypeEnum struct {
	SUB_EVENT_OPS    ListEventsRequestSubEventType
	SUB_EVENT_PLAN   ListEventsRequestSubEventType
	SUB_EVENT_CUSTOM ListEventsRequestSubEventType
}

func GetListEventsRequestSubEventTypeEnum() ListEventsRequestSubEventTypeEnum {
	return ListEventsRequestSubEventTypeEnum{
		SUB_EVENT_OPS: ListEventsRequestSubEventType{
			value: "SUB_EVENT.OPS",
		},
		SUB_EVENT_PLAN: ListEventsRequestSubEventType{
			value: "SUB_EVENT.PLAN",
		},
		SUB_EVENT_CUSTOM: ListEventsRequestSubEventType{
			value: "SUB_EVENT.CUSTOM",
		},
	}
}

func (c ListEventsRequestSubEventType) Value() string {
	return c.value
}

func (c ListEventsRequestSubEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventsRequestSubEventType) UnmarshalJSON(b []byte) error {
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
