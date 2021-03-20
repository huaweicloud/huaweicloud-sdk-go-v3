package model

import (
	"encoding/json"

	"strings"
)

// 转发kafka消息结构
type ActionKafkaForwarding struct {
	// 转发kafka消息对应的region区域

	RegionName *string `json:"region_name,omitempty"`
	// 转发kafka消息对应的projectId信息

	ProjectId *string `json:"project_id,omitempty"`
	// 转发kafka消息对应的地址列表

	KafkaAddresses *[]NetAddress `json:"kafka_addresses,omitempty"`
	// 转发kafka消息关联的topic信息。

	KafkaTopic *string `json:"kafka_topic,omitempty"`
	// 转发kafka关联的用户名信息。

	KafkaUsername *string `json:"kafka_username,omitempty"`
	// 转发kafka关联的密码信息。

	KafkaPassword *string `json:"kafka_password,omitempty"`
	// 转发kafka关联的鉴权机制。 类型说明： PAAS：非SASL鉴权。 PLAIN：SASL/PLAIN模式。需要填写对应的用户名密码信息。

	KafkaMechanism *string `json:"kafka_mechanism,omitempty"`
}

func (o ActionKafkaForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ActionKafkaForwarding struct{}"
	}

	return strings.Join([]string{"ActionKafkaForwarding", string(data)}, " ")
}
