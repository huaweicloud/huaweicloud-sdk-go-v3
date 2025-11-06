package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEventDetailResponse Response Object
type ListEventDetailResponse struct {

	// **参数解释**： 事件名称，值为系统产生的事件名称，或用户自定义上报的事件名称。 **取值范围**： 不涉及。
	EventName *string `json:"event_name,omitempty"`

	// **参数解释**： 事件类型。 **取值范围**： 值为EVENT.SYS或EVENT.CUSTOM。 - EVENT.SYS: 系统事件。 - EVENT.CUSTOM: 自定义事件。
	EventType *ListEventDetailResponseEventType `json:"event_type,omitempty"`

	// **参数解释**： 事件子类。 **取值范围**： 枚举类型。 当事件类型为系统事件时，参数值为SUB_EVENT.OPS或SUB_EVENT.PLAN。 当事件类型为自定义事件时，参数值为SUB_EVENT.CUSTOM。 - SUB_EVENT.OPS：运维事件。 - SUB_EVENT.PLAN：计划事件。 - SUB_EVENT.CUSTOM：自定义事件。
	SubEventType *ListEventDetailResponseSubEventType `json:"sub_event_type,omitempty"`

	// **参数解释**： 上报事件时用户的名称，也可能为projectID。 **取值范围**： 不涉及。
	EventUsers *[]string `json:"event_users,omitempty"`

	// **参数解释**： 事件来源。 如果是系统事件则值为各服务的命名空间，可查看支持监控的服务列表。如果是自定义事件，则为用户自定义上报定义。 **取值范围**： 不涉及。
	EventSources *[]string `json:"event_sources,omitempty"`

	// **参数解释**： 一条或者多条事件详细信息。
	EventInfo *[]EventInfoDetailResp `json:"event_info,omitempty"`

	MetaData       *TotalMetaData `json:"meta_data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListEventDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventDetailResponse struct{}"
	}

	return strings.Join([]string{"ListEventDetailResponse", string(data)}, " ")
}

type ListEventDetailResponseEventType struct {
	value string
}

type ListEventDetailResponseEventTypeEnum struct {
	EVENT_SYS    ListEventDetailResponseEventType
	EVENT_CUSTOM ListEventDetailResponseEventType
}

func GetListEventDetailResponseEventTypeEnum() ListEventDetailResponseEventTypeEnum {
	return ListEventDetailResponseEventTypeEnum{
		EVENT_SYS: ListEventDetailResponseEventType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: ListEventDetailResponseEventType{
			value: "EVENT.CUSTOM",
		},
	}
}

func (c ListEventDetailResponseEventType) Value() string {
	return c.value
}

func (c ListEventDetailResponseEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventDetailResponseEventType) UnmarshalJSON(b []byte) error {
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

type ListEventDetailResponseSubEventType struct {
	value string
}

type ListEventDetailResponseSubEventTypeEnum struct {
	SUB_EVENT_OPS    ListEventDetailResponseSubEventType
	SUB_EVENT_PLAN   ListEventDetailResponseSubEventType
	SUB_EVENT_CUSTOM ListEventDetailResponseSubEventType
}

func GetListEventDetailResponseSubEventTypeEnum() ListEventDetailResponseSubEventTypeEnum {
	return ListEventDetailResponseSubEventTypeEnum{
		SUB_EVENT_OPS: ListEventDetailResponseSubEventType{
			value: "SUB_EVENT.OPS",
		},
		SUB_EVENT_PLAN: ListEventDetailResponseSubEventType{
			value: "SUB_EVENT.PLAN",
		},
		SUB_EVENT_CUSTOM: ListEventDetailResponseSubEventType{
			value: "SUB_EVENT.CUSTOM",
		},
	}
}

func (c ListEventDetailResponseSubEventType) Value() string {
	return c.value
}

func (c ListEventDetailResponseSubEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventDetailResponseSubEventType) UnmarshalJSON(b []byte) error {
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
