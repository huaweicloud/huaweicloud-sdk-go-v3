package model

import (
	"encoding/json"

	"strings"
)

// 转发kafka消息结构
type DmsKafkaForwarding struct {
	// Kafka服务对应的region区域

	RegionName string `json:"region_name"`
	// Kafka服务对应的projectId信息

	ProjectId string `json:"project_id"`
	// 转发kafka消息对应的地址列表

	Addresses []NetAddress `json:"addresses"`
	// 转发kafka消息关联的topic信息。

	Topic string `json:"topic"`
	// 转发kafka关联的用户名信息。

	Username *string `json:"username,omitempty"`
	// 转发kafka关联的密码信息。

	Password *string `json:"password,omitempty"`
	// 转发kafka关联的鉴权机制。 类型说明： PAAS：非SASL鉴权。 PLAIN：SASL/PLAIN模式。需要填写对应的用户名密码信息。

	Mechanism *string `json:"mechanism,omitempty"`
}

func (o DmsKafkaForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DmsKafkaForwarding struct{}"
	}

	return strings.Join([]string{"DmsKafkaForwarding", string(data)}, " ")
}
