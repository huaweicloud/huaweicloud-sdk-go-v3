package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EventItemDetail
type EventItemDetail struct {

	// **参数解释**： 事件内容 **约束限制**： 不涉及。 **取值范围**： 长度为[0,4096]个字符。 **默认取值**： 不涉及。
	Content *string `json:"content,omitempty"`

	// **参数解释**： 所属分组。 资源分组对应的ID，必须是已存在的分组ID。 分组ID查询方法： 1.登录管理控制台。 2.单击“云监控服务”。 3.单击页面左侧的“资源分组”。 在名称/ID列获取具体资源分组ID。 **约束限制**： 不涉及。 **取值范围**： 长度只能为24个字符。 **默认取值**： 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 资源ID。 资源ID的查询方法： 1.登录管理控制台。 2.单击“计算 > 弹性云服务器”。 在资源概览页可获取具体资源ID。 **约束限制**： 不涉及。 **取值范围**： 支持字母、数字支持字母、数字、下划线（_）、中划线（-）和冒号（:），最大长度128个字符。例如，6a69bf28-ee62-49f3-9785-845dacd799ec。 **默认取值**： 不涉及。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 资源名称。 **约束限制**： 不涉及。 **取值范围**： 支持字母 中文 数字_ -. ，最大长度128个字符。 **默认取值**： 不涉及。
	ResourceName *string `json:"resource_name,omitempty"`

	// **参数解释**： 事件状态。 **约束限制**： 不涉及。 **取值范围**： 枚举类型。 - normal：正常发生 - warning：异常 - incident：严重 **默认取值**： 不涉及。
	EventState *EventItemDetailEventState `json:"event_state,omitempty"`

	// **参数解释**： 事件级别。 **约束限制**： 不涉及。 **取值范围**： 枚举类型：Critical, Major, Minor, Info。 - Critical: 紧急 - Major: 重要 - Minor: 次要 - Info: 提示 **默认取值**： 不涉及。
	EventLevel *EventItemDetailEventLevel `json:"event_level,omitempty"`

	// **参数解释**： 事件用户。 **约束限制**： 不涉及。 **取值范围**： 支持字母 数字_ -/空格 ，长度为[0,64]个字符。 **默认取值**： 不涉及。
	EventUser *string `json:"event_user,omitempty"`

	// **参数解释**： 事件类型。 **约束限制**： EVENT.SYS为系统事件，用户自己不能上报系统事件，只能传EVENT.CUSTOM。 **取值范围**： 枚举类型，EVENT.SYS或EVENT.CUSTOM。 - EVENT.SYS：系统事件 - EVENT.CUSTOM：自定义事件 **默认取值**： 不涉及。
	EventType *EventItemDetailEventType `json:"event_type,omitempty"`

	// **参数解释**： 事件子类。 **约束限制**： 不涉及。 **取值范围**： 枚举类型 - SUB_EVENT.OPS: 运维事件 - SUB_EVENT.PLAN: 计划事件 - SUB_EVENT.CUSTOM: 自定义事件 **默认取值**： 不涉及。
	SubEventType *EventItemDetailSubEventType `json:"sub_event_type,omitempty"`

	// **参数解释**： 事件的维度，根据维度描述资源信息。 用于指定资源、资源分组的事件告警场景中，支持按维度配置告警规则。 **约束限制**： 目前最大支持4个维度。
	Dimensions *[]MetricsDimension `json:"dimensions,omitempty"`
}

func (o EventItemDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventItemDetail struct{}"
	}

	return strings.Join([]string{"EventItemDetail", string(data)}, " ")
}

type EventItemDetailEventState struct {
	value string
}

type EventItemDetailEventStateEnum struct {
	NORMAL   EventItemDetailEventState
	WARNING  EventItemDetailEventState
	INCIDENT EventItemDetailEventState
}

func GetEventItemDetailEventStateEnum() EventItemDetailEventStateEnum {
	return EventItemDetailEventStateEnum{
		NORMAL: EventItemDetailEventState{
			value: "normal",
		},
		WARNING: EventItemDetailEventState{
			value: "warning",
		},
		INCIDENT: EventItemDetailEventState{
			value: "incident",
		},
	}
}

func (c EventItemDetailEventState) Value() string {
	return c.value
}

func (c EventItemDetailEventState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventItemDetailEventState) UnmarshalJSON(b []byte) error {
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

type EventItemDetailEventLevel struct {
	value string
}

type EventItemDetailEventLevelEnum struct {
	CRITICAL EventItemDetailEventLevel
	MAJOR    EventItemDetailEventLevel
	MINOR    EventItemDetailEventLevel
	INFO     EventItemDetailEventLevel
}

func GetEventItemDetailEventLevelEnum() EventItemDetailEventLevelEnum {
	return EventItemDetailEventLevelEnum{
		CRITICAL: EventItemDetailEventLevel{
			value: "Critical",
		},
		MAJOR: EventItemDetailEventLevel{
			value: "Major",
		},
		MINOR: EventItemDetailEventLevel{
			value: "Minor",
		},
		INFO: EventItemDetailEventLevel{
			value: "Info",
		},
	}
}

func (c EventItemDetailEventLevel) Value() string {
	return c.value
}

func (c EventItemDetailEventLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventItemDetailEventLevel) UnmarshalJSON(b []byte) error {
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

type EventItemDetailEventType struct {
	value string
}

type EventItemDetailEventTypeEnum struct {
	EVENT_SYS    EventItemDetailEventType
	EVENT_CUSTOM EventItemDetailEventType
}

func GetEventItemDetailEventTypeEnum() EventItemDetailEventTypeEnum {
	return EventItemDetailEventTypeEnum{
		EVENT_SYS: EventItemDetailEventType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: EventItemDetailEventType{
			value: "EVENT.CUSTOM",
		},
	}
}

func (c EventItemDetailEventType) Value() string {
	return c.value
}

func (c EventItemDetailEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventItemDetailEventType) UnmarshalJSON(b []byte) error {
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

type EventItemDetailSubEventType struct {
	value string
}

type EventItemDetailSubEventTypeEnum struct {
	SUB_EVENT_OPS    EventItemDetailSubEventType
	SUB_EVENT_PLAN   EventItemDetailSubEventType
	SUB_EVENT_CUSTOM EventItemDetailSubEventType
}

func GetEventItemDetailSubEventTypeEnum() EventItemDetailSubEventTypeEnum {
	return EventItemDetailSubEventTypeEnum{
		SUB_EVENT_OPS: EventItemDetailSubEventType{
			value: "SUB_EVENT.OPS",
		},
		SUB_EVENT_PLAN: EventItemDetailSubEventType{
			value: "SUB_EVENT.PLAN",
		},
		SUB_EVENT_CUSTOM: EventItemDetailSubEventType{
			value: "SUB_EVENT.CUSTOM",
		},
	}
}

func (c EventItemDetailSubEventType) Value() string {
	return c.value
}

func (c EventItemDetailSubEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventItemDetailSubEventType) UnmarshalJSON(b []byte) error {
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
