package model

import (
	"encoding/json"

	"strings"
)

// amqp queue配置信息
type AmqpForwarding struct {
	// **参数说明**：用于接收满足规则条件数据的amqp queue。

	QueueName string `json:"queue_name"`
}

func (o AmqpForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AmqpForwarding struct{}"
	}

	return strings.Join([]string{"AmqpForwarding", string(data)}, " ")
}
