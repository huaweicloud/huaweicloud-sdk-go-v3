package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EventStreamingOperateReq struct {

	// 操作类型
	Operation *EventStreamingOperateReqOperation `json:"operation,omitempty"`
}

func (o EventStreamingOperateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventStreamingOperateReq struct{}"
	}

	return strings.Join([]string{"EventStreamingOperateReq", string(data)}, " ")
}

type EventStreamingOperateReqOperation struct {
	value string
}

type EventStreamingOperateReqOperationEnum struct {
	START EventStreamingOperateReqOperation
	PAUSE EventStreamingOperateReqOperation
}

func GetEventStreamingOperateReqOperationEnum() EventStreamingOperateReqOperationEnum {
	return EventStreamingOperateReqOperationEnum{
		START: EventStreamingOperateReqOperation{
			value: "START",
		},
		PAUSE: EventStreamingOperateReqOperation{
			value: "PAUSE",
		},
	}
}

func (c EventStreamingOperateReqOperation) Value() string {
	return c.value
}

func (c EventStreamingOperateReqOperation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventStreamingOperateReqOperation) UnmarshalJSON(b []byte) error {
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
