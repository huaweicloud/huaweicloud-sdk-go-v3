package model

import (
	"encoding/json"

	"strings"
)

// pulsar通道配置信
type PulsarForwarding struct {
	// **参数说明**：pulsar的访问url。

	Url string `json:"url"`
	// **参数说明**：用于接收满足规则条件数据的topic。

	Topic string `json:"topic"`
	// **参数说明**：证书id

	CertId *string `json:"cert_id,omitempty"`
}

func (o PulsarForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PulsarForwarding struct{}"
	}

	return strings.Join([]string{"PulsarForwarding", string(data)}, " ")
}
