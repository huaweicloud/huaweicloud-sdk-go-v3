package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EventStreamingSource 事件源，一个事件流中事件源只有一个
type EventStreamingSource struct {
	SourceKafka *SourceKafkaMqParameters `json:"source_kafka,omitempty"`

	// 事件源类型名称
	Name *EventStreamingSourceName `json:"name,omitempty"`
}

func (o EventStreamingSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventStreamingSource struct{}"
	}

	return strings.Join([]string{"EventStreamingSource", string(data)}, " ")
}

type EventStreamingSourceName struct {
	value string
}

type EventStreamingSourceNameEnum struct {
	HC_DMS EventStreamingSourceName
}

func GetEventStreamingSourceNameEnum() EventStreamingSourceNameEnum {
	return EventStreamingSourceNameEnum{
		HC_DMS: EventStreamingSourceName{
			value: "HC.DMS",
		},
	}
}

func (c EventStreamingSourceName) Value() string {
	return c.value
}

func (c EventStreamingSourceName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventStreamingSourceName) UnmarshalJSON(b []byte) error {
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
