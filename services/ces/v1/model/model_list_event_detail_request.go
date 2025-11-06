package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEventDetailRequest Request Object
type ListEventDetailRequest struct {

	// **参数解释**： 事件名称，值为系统产生的事件名称或用户自定义上报的事件名称。 **约束限制**： 不涉及。 **取值范围**： 长度为[1,64]个字符。 **默认取值**： 不涉及。
	EventName string `json:"event_name"`

	// **参数解释**： 事件类型。 **约束限制**： 不涉及。 **取值范围**： 值为EVENT.SYS或EVENT.CUSTOM。 - EVENT.SYS：系统事件。 - EVENT.CUSTOM：自定义事件。 **默认取值**： 不涉及。
	EventType ListEventDetailRequestEventType `json:"event_type"`

	// **参数解释**： 事件子类。 **约束限制**： 不涉及。 **取值范围**： 枚举类型 - SUB_EVENT.OPS: 运维事件 - SUB_EVENT.PLAN: 计划事件 - SUB_EVENT.CUSTOM: 自定义事件 **默认取值**： 不涉及。
	SubEventType *ListEventDetailRequestSubEventType `json:"sub_event_type,omitempty"`

	// **参数解释**： 事件来源，取值为各云服务的命名空间。云服务的命名空间请参考“[支持监控的服务列表](ces_03_0059.xml)”。 **约束限制**： 不涉及。 **取值范围**： 长度为[0,32]个字符。 正则匹配：以字母开头，中间有一个点，包含字母、数字、下划线的字符串。 **默认取值**： 不涉及。
	EventSource *string `json:"event_source,omitempty"`

	// **参数解释**： 事件的级别。 **约束限制**： 不涉及。 **取值范围**： 值为Critical、Major、Minor、Info。 - Critical: 紧急 - Major: 重要 - Minor: 次要 - Info: 提示 **默认取值**： 不涉及。
	EventLevel *ListEventDetailRequestEventLevel `json:"event_level,omitempty"`

	// **参数解释**： 上报事件监控数据时用户的名称，也可为projectID。 **约束限制**： 不涉及。 **取值范围**： 长度为[0,64]个字符。 正则匹配：由零个或多个字母、数字、下划线、横线、斜杠、空格、@ 符号或点号组成的字符串。 **默认取值**： 不涉及。
	EventUser *string `json:"event_user,omitempty"`

	// **参数解释**： 事件的状态。 **约束限制**： 不涉及。 **取值范围**： 值为normal、warning、incident。 - normal: 正常 - warning: 警告 - incident: 故障 **默认取值**： 不涉及。
	EventState *ListEventDetailRequestEventState `json:"event_state,omitempty"`

	// **参数解释**： 查询数据起始时间，UNIX时间戳，单位毫秒。例如：1605952700911。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	From *int64 `json:"from,omitempty"`

	// **参数解释**： 查询数据截止时间，UNIX时间戳，单位毫秒。 **约束限制**： 其中from必须小于to。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	To *int64 `json:"to,omitempty"`

	// **参数解释**： 分页起始值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	Start *int64 `json:"start,omitempty"`

	// **参数解释**： 单次查询的条数限制，用于限制结果数据条数。 **约束限制**： 不涉及。 **取值范围**： 大小为[1,100]的整数 **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEventDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventDetailRequest struct{}"
	}

	return strings.Join([]string{"ListEventDetailRequest", string(data)}, " ")
}

type ListEventDetailRequestEventType struct {
	value string
}

type ListEventDetailRequestEventTypeEnum struct {
	EVENT_SYS    ListEventDetailRequestEventType
	EVENT_CUSTOM ListEventDetailRequestEventType
}

func GetListEventDetailRequestEventTypeEnum() ListEventDetailRequestEventTypeEnum {
	return ListEventDetailRequestEventTypeEnum{
		EVENT_SYS: ListEventDetailRequestEventType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: ListEventDetailRequestEventType{
			value: "EVENT.CUSTOM",
		},
	}
}

func (c ListEventDetailRequestEventType) Value() string {
	return c.value
}

func (c ListEventDetailRequestEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventDetailRequestEventType) UnmarshalJSON(b []byte) error {
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

type ListEventDetailRequestSubEventType struct {
	value string
}

type ListEventDetailRequestSubEventTypeEnum struct {
	SUB_EVENT_OPS    ListEventDetailRequestSubEventType
	SUB_EVENT_PLAN   ListEventDetailRequestSubEventType
	SUB_EVENT_CUSTOM ListEventDetailRequestSubEventType
}

func GetListEventDetailRequestSubEventTypeEnum() ListEventDetailRequestSubEventTypeEnum {
	return ListEventDetailRequestSubEventTypeEnum{
		SUB_EVENT_OPS: ListEventDetailRequestSubEventType{
			value: "SUB_EVENT.OPS",
		},
		SUB_EVENT_PLAN: ListEventDetailRequestSubEventType{
			value: "SUB_EVENT.PLAN",
		},
		SUB_EVENT_CUSTOM: ListEventDetailRequestSubEventType{
			value: "SUB_EVENT.CUSTOM",
		},
	}
}

func (c ListEventDetailRequestSubEventType) Value() string {
	return c.value
}

func (c ListEventDetailRequestSubEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventDetailRequestSubEventType) UnmarshalJSON(b []byte) error {
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

type ListEventDetailRequestEventLevel struct {
	value string
}

type ListEventDetailRequestEventLevelEnum struct {
	CRITICAL ListEventDetailRequestEventLevel
	MAJOR    ListEventDetailRequestEventLevel
	INFO     ListEventDetailRequestEventLevel
	MINOR    ListEventDetailRequestEventLevel
}

func GetListEventDetailRequestEventLevelEnum() ListEventDetailRequestEventLevelEnum {
	return ListEventDetailRequestEventLevelEnum{
		CRITICAL: ListEventDetailRequestEventLevel{
			value: "Critical",
		},
		MAJOR: ListEventDetailRequestEventLevel{
			value: "Major",
		},
		INFO: ListEventDetailRequestEventLevel{
			value: "Info",
		},
		MINOR: ListEventDetailRequestEventLevel{
			value: "Minor",
		},
	}
}

func (c ListEventDetailRequestEventLevel) Value() string {
	return c.value
}

func (c ListEventDetailRequestEventLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventDetailRequestEventLevel) UnmarshalJSON(b []byte) error {
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

type ListEventDetailRequestEventState struct {
	value string
}

type ListEventDetailRequestEventStateEnum struct {
	NORMAL   ListEventDetailRequestEventState
	WARNING  ListEventDetailRequestEventState
	INCIDENT ListEventDetailRequestEventState
}

func GetListEventDetailRequestEventStateEnum() ListEventDetailRequestEventStateEnum {
	return ListEventDetailRequestEventStateEnum{
		NORMAL: ListEventDetailRequestEventState{
			value: "normal",
		},
		WARNING: ListEventDetailRequestEventState{
			value: "warning",
		},
		INCIDENT: ListEventDetailRequestEventState{
			value: "incident",
		},
	}
}

func (c ListEventDetailRequestEventState) Value() string {
	return c.value
}

func (c ListEventDetailRequestEventState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventDetailRequestEventState) UnmarshalJSON(b []byte) error {
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
