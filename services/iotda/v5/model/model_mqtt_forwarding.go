package model

import (
	"encoding/json"

	"strings"
)

// topic配置信息
type MqttForwarding struct {
	// 用于接收满足规则条件数据的topic。
	Topic string `json:"topic"`
}

func (o MqttForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MqttForwarding struct{}"
	}

	return strings.Join([]string{"MqttForwarding", string(data)}, " ")
}
