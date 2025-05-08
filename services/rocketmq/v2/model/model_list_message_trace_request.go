package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMessageTraceRequest Request Object
type ListMessageTraceRequest struct {

	// 消息引擎。
	Engine ListMessageTraceRequestEngine `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消息ID。
	MsgId string `json:"msg_id"`

	// 查询数量。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`
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

func (c ListMessageTraceRequestEngine) Value() string {
	return c.value
}

func (c ListMessageTraceRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMessageTraceRequestEngine) UnmarshalJSON(b []byte) error {
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
