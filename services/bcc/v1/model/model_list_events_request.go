package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEventsRequest Request Object
type ListEventsRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 事件源，取值范围：SYS.CBR,SYS.RDS,SYS.GaussDB
	EventSource *ListEventsRequestEventSource `json:"event_source,omitempty"`

	// 事件等级，取值范围：Critical,Major,Minor,Info
	EventLevel *ListEventsRequestEventLevel `json:"event_level,omitempty"`

	// 事件名称，长度范围4-64个字符。
	EventName *string `json:"event_name,omitempty"`

	// 资源ID，长度范围16-64个字符。
	ResourceId *string `json:"resource_id,omitempty"`

	// 查询范围起始时间，例如：2025-03-20T09:31:45Z。
	StartTime *string `json:"start_time,omitempty"`

	// 查询范围结束时间，例如：2025-03-20T09:31:45Z。
	EndTime *string `json:"end_time,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页。
	Marker *string `json:"marker,omitempty"`

	// 单次查询的条数限制,取值范围(0,100]，默认值为10，用于限制结果数据条数。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsRequest struct{}"
	}

	return strings.Join([]string{"ListEventsRequest", string(data)}, " ")
}

type ListEventsRequestEventSource struct {
	value string
}

type ListEventsRequestEventSourceEnum struct {
	SYS_CBR      ListEventsRequestEventSource
	SYS_RDS      ListEventsRequestEventSource
	SYS_GAUSS_DB ListEventsRequestEventSource
}

func GetListEventsRequestEventSourceEnum() ListEventsRequestEventSourceEnum {
	return ListEventsRequestEventSourceEnum{
		SYS_CBR: ListEventsRequestEventSource{
			value: "SYS.CBR",
		},
		SYS_RDS: ListEventsRequestEventSource{
			value: "SYS.RDS",
		},
		SYS_GAUSS_DB: ListEventsRequestEventSource{
			value: "SYS.GaussDB",
		},
	}
}

func (c ListEventsRequestEventSource) Value() string {
	return c.value
}

func (c ListEventsRequestEventSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventsRequestEventSource) UnmarshalJSON(b []byte) error {
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

type ListEventsRequestEventLevel struct {
	value string
}

type ListEventsRequestEventLevelEnum struct {
	CRITICAL ListEventsRequestEventLevel
	MAJOR    ListEventsRequestEventLevel
	MINOR    ListEventsRequestEventLevel
	INFO     ListEventsRequestEventLevel
}

func GetListEventsRequestEventLevelEnum() ListEventsRequestEventLevelEnum {
	return ListEventsRequestEventLevelEnum{
		CRITICAL: ListEventsRequestEventLevel{
			value: "Critical",
		},
		MAJOR: ListEventsRequestEventLevel{
			value: "Major",
		},
		MINOR: ListEventsRequestEventLevel{
			value: "Minor",
		},
		INFO: ListEventsRequestEventLevel{
			value: "Info",
		},
	}
}

func (c ListEventsRequestEventLevel) Value() string {
	return c.value
}

func (c ListEventsRequestEventLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventsRequestEventLevel) UnmarshalJSON(b []byte) error {
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
