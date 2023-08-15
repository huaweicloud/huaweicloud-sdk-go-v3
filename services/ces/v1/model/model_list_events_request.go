package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEventsRequest Request Object
type ListEventsRequest struct {

	// 事件类型，值为EVENT.SYS或EVENT.CUSTOM，EVENT.SYS表示系统事件，EVENT.CUSTOM表示自定义事件。
	EventType *ListEventsRequestEventType `json:"event_type,omitempty"`

	// 事件名称，值为系统产生的事件名称，或用户自定义上报的事件名称。
	EventName *string `json:"event_name,omitempty"`

	// 查询数据起始时间，UNIX时间戳，单位毫秒；例如：1605952700911。
	From *int64 `json:"from,omitempty"`

	// 查询数据截止时间UNIX时间戳，单位毫秒。from必须小于to，例如：1606557500911。
	To *int64 `json:"to,omitempty"`

	// 分页起始值，类型为integer，默认值为0。
	Start *int32 `json:"start,omitempty"`

	// 单次查询的条数限制，取值范围(0,100]，默认值为100，用于限制结果数据条数。
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
