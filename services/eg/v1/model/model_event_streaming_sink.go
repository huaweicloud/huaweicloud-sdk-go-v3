package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EventStreamingSink 事件目标，一个事件流中只有一个事件目标，sink_fg、sink_kafka值能选择其中一个参数
type EventStreamingSink struct {
	SinkFg *SinkFgParameters `json:"sink_fg,omitempty"`

	SinkKafka *SinkKafkaParameters `json:"sink_kafka,omitempty"`

	// 事件目标类型名称
	Name *EventStreamingSinkName `json:"name,omitempty"`
}

func (o EventStreamingSink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventStreamingSink struct{}"
	}

	return strings.Join([]string{"EventStreamingSink", string(data)}, " ")
}

type EventStreamingSinkName struct {
	value string
}

type EventStreamingSinkNameEnum struct {
	HC_FUNCTION_GRAPH EventStreamingSinkName
	HC_KAFKA          EventStreamingSinkName
}

func GetEventStreamingSinkNameEnum() EventStreamingSinkNameEnum {
	return EventStreamingSinkNameEnum{
		HC_FUNCTION_GRAPH: EventStreamingSinkName{
			value: "HC.FunctionGraph",
		},
		HC_KAFKA: EventStreamingSinkName{
			value: "HC.Kafka",
		},
	}
}

func (c EventStreamingSinkName) Value() string {
	return c.value
}

func (c EventStreamingSinkName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventStreamingSinkName) UnmarshalJSON(b []byte) error {
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
