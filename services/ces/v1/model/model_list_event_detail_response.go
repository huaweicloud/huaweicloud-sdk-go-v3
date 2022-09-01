package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ListEventDetailResponse struct {

	// 事件名称，值为系统产生的事件名称，或用户自定义上报的事件名称。
	EventName *string `json:"event_name,omitempty" xml:"event_name"`

	// 事件类型，值为EVENT.SYS或EVENT.CUSTOM，EVENT.SYS表示系统事件，EVENT.CUSTOM表示自定义事件。
	EventType *ListEventDetailResponseEventType `json:"event_type,omitempty" xml:"event_type"`

	// 上报事件时用户的名称，也可能为projectID。
	EventUsers *[]string `json:"event_users,omitempty" xml:"event_users"`

	// 事件来源，如果是系统事件则值为各服务的命名空间，各服务的命名空间可查看：“[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”；如果是自定义事件，则为用户自定义上报定义。
	EventSources *[]string `json:"event_sources,omitempty" xml:"event_sources"`

	// 一条或者多条事件详细信息。
	EventInfo *[]EventInfoDetail `json:"event_info,omitempty" xml:"event_info"`

	MetaData       *TotalMetaData `json:"meta_data,omitempty" xml:"meta_data"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
