package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeadLetterQueue 订阅的死信参数列表
type DeadLetterQueue struct {

	// 队列类型
	Type DeadLetterQueueType `json:"type"`

	// 实例id
	InstanceId string `json:"instance_id"`

	// 目标连接id
	ConnectionId string `json:"connection_id"`

	// 主题
	Topic string `json:"topic"`
}

func (o DeadLetterQueue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeadLetterQueue struct{}"
	}

	return strings.Join([]string{"DeadLetterQueue", string(data)}, " ")
}

type DeadLetterQueueType struct {
	value string
}

type DeadLetterQueueTypeEnum struct {
	KAFKA DeadLetterQueueType
}

func GetDeadLetterQueueTypeEnum() DeadLetterQueueTypeEnum {
	return DeadLetterQueueTypeEnum{
		KAFKA: DeadLetterQueueType{
			value: "KAFKA",
		},
	}
}

func (c DeadLetterQueueType) Value() string {
	return c.value
}

func (c DeadLetterQueueType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeadLetterQueueType) UnmarshalJSON(b []byte) error {
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
