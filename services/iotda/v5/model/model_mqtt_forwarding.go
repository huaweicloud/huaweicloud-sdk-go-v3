package model

import (
	"encoding/json"

	"strings"
)

// topic配置信息
type MqttForwarding struct {
	// **参数说明**：用于接收满足规则条件数据的topic。 **取值范围**：长度不超过128，只允许字母、数字、下划线（_）、斜杠（/）、连接符（-）的组合。

	Topic string `json:"topic"`
}

func (o MqttForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MqttForwarding struct{}"
	}

	return strings.Join([]string{"MqttForwarding", string(data)}, " ")
}
