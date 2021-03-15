package model

import (
	"encoding/json"

	"strings"
)

// 转发MRS Kafka消息结构
type MrsKafkaForwarding struct {
	// 转发kafka消息对应的地址列表
	Addresses []NetAddress `json:"addresses"`
	// 转发kafka消息关联的topic信息。
	Topic string `json:"topic"`
}

func (o MrsKafkaForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MrsKafkaForwarding struct{}"
	}

	return strings.Join([]string{"MrsKafkaForwarding", string(data)}, " ")
}
