package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEventItemDetailResp
type ShowEventItemDetailResp struct {

	// **参数解释**： 事件内容。 **取值范围**： 长度为[1,4096]个字符。
	Content *string `json:"content,omitempty"`

	// **参数解释**： 所属分组。 **取值范围**： 长度只能为24个字符。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 资源ID。 **取值范围**： 长度为[1,128]个字符。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 资源名称。 **取值范围**： 长度为[1,128]个字符。
	ResourceName *string `json:"resource_name,omitempty"`

	// **参数解释**： 事件状态。 **取值范围**： 枚举类型：normal\\warning\\incident。 - normal: 正常 - warning: 警告 - incident: 故障
	EventState *ShowEventItemDetailRespEventState `json:"event_state,omitempty"`

	// **参数解释**： 事件级别。 **取值范围**： 枚举类型：Critical, Major, Minor, Info。 - Critical: 紧急 - Major: 重要 - Minor: 次要 - Info: 提示
	EventLevel *ShowEventItemDetailRespEventLevel `json:"event_level,omitempty"`

	// **参数解释**： 事件用户。 **取值范围**： 下划线、横线、斜杠、@ 符号或点号组成，长度为[0,64]个字符。
	EventUser *string `json:"event_user,omitempty"`

	// **参数解释**： 事件类型。 **取值范围**： 枚举类型：EVENT.SYS，EVENT.CUSTOM - EVENT.SYS: 系统事件。 - EVENT.CUSTOM: 自定义事件。
	EventType *string `json:"event_type,omitempty"`

	// **参数解释**： 事件子类。 **取值范围**： 枚举类型。 当事件类型为系统事件时，参数值为SUB_EVENT.OPS或SUB_EVENT.PLAN。 当事件类型为自定义事件时，参数值为SUB_EVENT.CUSTOM。 - SUB_EVENT.OPS：运维事件。 - SUB_EVENT.PLAN：计划事件。 - SUB_EVENT.CUSTOM：自定义事件。
	SubEventType *ShowEventItemDetailRespSubEventType `json:"sub_event_type,omitempty"`

	// **参数解释**： 事件的维度，根据维度描述资源信息。 用于指定资源、资源分组的事件告警场景中，支持按维度配置告警规则。 **取值范围**： 目前最大支持4个维度
	Dimensions *[]MetricsDimensionResp `json:"dimensions,omitempty"`
}

func (o ShowEventItemDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventItemDetailResp struct{}"
	}

	return strings.Join([]string{"ShowEventItemDetailResp", string(data)}, " ")
}

type ShowEventItemDetailRespEventState struct {
	value string
}

type ShowEventItemDetailRespEventStateEnum struct {
	NORMAL   ShowEventItemDetailRespEventState
	WARNING  ShowEventItemDetailRespEventState
	INCIDENT ShowEventItemDetailRespEventState
}

func GetShowEventItemDetailRespEventStateEnum() ShowEventItemDetailRespEventStateEnum {
	return ShowEventItemDetailRespEventStateEnum{
		NORMAL: ShowEventItemDetailRespEventState{
			value: "normal",
		},
		WARNING: ShowEventItemDetailRespEventState{
			value: "warning",
		},
		INCIDENT: ShowEventItemDetailRespEventState{
			value: "incident",
		},
	}
}

func (c ShowEventItemDetailRespEventState) Value() string {
	return c.value
}

func (c ShowEventItemDetailRespEventState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEventItemDetailRespEventState) UnmarshalJSON(b []byte) error {
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

type ShowEventItemDetailRespEventLevel struct {
	value string
}

type ShowEventItemDetailRespEventLevelEnum struct {
	CRITICAL ShowEventItemDetailRespEventLevel
	MAJOR    ShowEventItemDetailRespEventLevel
	MINOR    ShowEventItemDetailRespEventLevel
	INFO     ShowEventItemDetailRespEventLevel
}

func GetShowEventItemDetailRespEventLevelEnum() ShowEventItemDetailRespEventLevelEnum {
	return ShowEventItemDetailRespEventLevelEnum{
		CRITICAL: ShowEventItemDetailRespEventLevel{
			value: "Critical",
		},
		MAJOR: ShowEventItemDetailRespEventLevel{
			value: "Major",
		},
		MINOR: ShowEventItemDetailRespEventLevel{
			value: "Minor",
		},
		INFO: ShowEventItemDetailRespEventLevel{
			value: "Info",
		},
	}
}

func (c ShowEventItemDetailRespEventLevel) Value() string {
	return c.value
}

func (c ShowEventItemDetailRespEventLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEventItemDetailRespEventLevel) UnmarshalJSON(b []byte) error {
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

type ShowEventItemDetailRespSubEventType struct {
	value string
}

type ShowEventItemDetailRespSubEventTypeEnum struct {
	SUB_EVENT_OPS    ShowEventItemDetailRespSubEventType
	SUB_EVENT_PLAN   ShowEventItemDetailRespSubEventType
	SUB_EVENT_CUSTOM ShowEventItemDetailRespSubEventType
}

func GetShowEventItemDetailRespSubEventTypeEnum() ShowEventItemDetailRespSubEventTypeEnum {
	return ShowEventItemDetailRespSubEventTypeEnum{
		SUB_EVENT_OPS: ShowEventItemDetailRespSubEventType{
			value: "SUB_EVENT.OPS",
		},
		SUB_EVENT_PLAN: ShowEventItemDetailRespSubEventType{
			value: "SUB_EVENT.PLAN",
		},
		SUB_EVENT_CUSTOM: ShowEventItemDetailRespSubEventType{
			value: "SUB_EVENT.CUSTOM",
		},
	}
}

func (c ShowEventItemDetailRespSubEventType) Value() string {
	return c.value
}

func (c ShowEventItemDetailRespSubEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEventItemDetailRespSubEventType) UnmarshalJSON(b []byte) error {
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
