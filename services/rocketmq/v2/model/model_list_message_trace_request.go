package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListMessageTraceRequest struct {

	// 消息引擎。
	Engine ListMessageTraceRequestEngine `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消息ID。
	MsgId *string `json:"msg_id,omitempty"`
}

func (o ListMessageTraceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageTraceRequest struct{}"
	}

	return strings.Join([]string{"ListMessageTraceRequest", string(data)}, " ")
}

type ListMessageTraceRequestEngine struct {
	value string
}

type ListMessageTraceRequestEngineEnum struct {
	RELIABILITY ListMessageTraceRequestEngine
}

func GetListMessageTraceRequestEngineEnum() ListMessageTraceRequestEngineEnum {
	return ListMessageTraceRequestEngineEnum{
		RELIABILITY: ListMessageTraceRequestEngine{
			value: "reliability",
		},
	}
}

func (c ListMessageTraceRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMessageTraceRequestEngine) UnmarshalJSON(b []byte) error {
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
