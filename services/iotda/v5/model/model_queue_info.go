package model

import (
	"encoding/json"

	"strings"
)

// 添加时队列结构体。
type QueueInfo struct {
	// 队列名称，同一租户不允许重复。

	QueueName string `json:"queue_name"`
}

func (o QueueInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueueInfo struct{}"
	}

	return strings.Join([]string{"QueueInfo", string(data)}, " ")
}
