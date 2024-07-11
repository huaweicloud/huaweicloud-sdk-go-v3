package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMessagesRequest Request Object
type ListMessagesRequest struct {

	// 消息引擎。
	Engine ListMessagesRequestEngine `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 主题名称。
	Topic string `json:"topic"`

	// 查询数量。
	Limit *string `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *string `json:"offset,omitempty"`

	// 消息的key。
	Key *string `json:"key,omitempty"`

	// 开始时间（不通过msg_id精确查询消息时，此参数必填）。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间（不通过msg_id精确查询消息时，此参数必填）。
	EndTime *string `json:"end_time,omitempty"`

	// 消息ID。
	MsgId *string `json:"msg_id,omitempty"`
}

func (o ListMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessagesRequest struct{}"
	}

	return strings.Join([]string{"ListMessagesRequest", string(data)}, " ")
}

type ListMessagesRequestEngine struct {
	value string
}

type ListMessagesRequestEngineEnum struct {
	RELIABILITY ListMessagesRequestEngine
}

func GetListMessagesRequestEngineEnum() ListMessagesRequestEngineEnum {
	return ListMessagesRequestEngineEnum{
		RELIABILITY: ListMessagesRequestEngine{
			value: "reliability",
		},
	}
}

func (c ListMessagesRequestEngine) Value() string {
	return c.value
}

func (c ListMessagesRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMessagesRequestEngine) UnmarshalJSON(b []byte) error {
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
