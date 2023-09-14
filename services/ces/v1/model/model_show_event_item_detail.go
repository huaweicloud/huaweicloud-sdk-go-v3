package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEventItemDetail
type ShowEventItemDetail struct {

	// 事件内容，最大长度4096。
	Content *string `json:"content,omitempty"`

	// 所属分组。  资源分组对应的ID，必须传存在的分组ID。
	GroupId *string `json:"group_id,omitempty"`

	// 资源ID，支持字母、数字_ -：，最大长度128。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称，支持字母 中文 数字_ -. ，最大长度128。
	ResourceName *string `json:"resource_name,omitempty"`

	// 事件状态。  枚举类型：normal\\warning\\incident
	EventState *ShowEventItemDetailEventState `json:"event_state,omitempty"`

	// 事件级别。  枚举类型：Critical, Major, Minor, Info
	EventLevel *ShowEventItemDetailEventLevel `json:"event_level,omitempty"`

	// 事件用户。  支持字母 数字_ -/空格 ，最大长度64。
	EventUser *string `json:"event_user,omitempty"`

	// 事件类型。 枚举类型，EVENT.SYS或EVENT.CUSTOM，EVENT.SYS为系统事件，用户自已不能上报，只能传EVENT.CUSTOM。
	EventType *string `json:"event_type,omitempty"`

	// 维度信息列表
	Dimensions *[]MetricsDimension `json:"dimensions,omitempty"`
}

func (o ShowEventItemDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventItemDetail struct{}"
	}

	return strings.Join([]string{"ShowEventItemDetail", string(data)}, " ")
}

type ShowEventItemDetailEventState struct {
	value string
}

type ShowEventItemDetailEventStateEnum struct {
	NORMAL   ShowEventItemDetailEventState
	WARNING  ShowEventItemDetailEventState
	INCIDENT ShowEventItemDetailEventState
}

func GetShowEventItemDetailEventStateEnum() ShowEventItemDetailEventStateEnum {
	return ShowEventItemDetailEventStateEnum{
		NORMAL: ShowEventItemDetailEventState{
			value: "normal",
		},
		WARNING: ShowEventItemDetailEventState{
			value: "warning",
		},
		INCIDENT: ShowEventItemDetailEventState{
			value: "incident",
		},
	}
}

func (c ShowEventItemDetailEventState) Value() string {
	return c.value
}

func (c ShowEventItemDetailEventState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEventItemDetailEventState) UnmarshalJSON(b []byte) error {
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

type ShowEventItemDetailEventLevel struct {
	value string
}

type ShowEventItemDetailEventLevelEnum struct {
	CRITICAL ShowEventItemDetailEventLevel
	MAJOR    ShowEventItemDetailEventLevel
	MINOR    ShowEventItemDetailEventLevel
	INFO     ShowEventItemDetailEventLevel
}

func GetShowEventItemDetailEventLevelEnum() ShowEventItemDetailEventLevelEnum {
	return ShowEventItemDetailEventLevelEnum{
		CRITICAL: ShowEventItemDetailEventLevel{
			value: "Critical",
		},
		MAJOR: ShowEventItemDetailEventLevel{
			value: "Major",
		},
		MINOR: ShowEventItemDetailEventLevel{
			value: "Minor",
		},
		INFO: ShowEventItemDetailEventLevel{
			value: "Info",
		},
	}
}

func (c ShowEventItemDetailEventLevel) Value() string {
	return c.value
}

func (c ShowEventItemDetailEventLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEventItemDetailEventLevel) UnmarshalJSON(b []byte) error {
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
