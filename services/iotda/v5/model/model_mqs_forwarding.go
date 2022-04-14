package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 服务配置信息
type MqsForwarding struct {
	// **参数说明**：MQS服务的URL

	Url string `json:"url"`
	// **参数说明**：用于登录MQS的用户名

	UserName string `json:"user_name"`
	// **参数说明**：用于登录MQS的密码

	Password string `json:"password"`
	// **参数说明**：订阅的MQS主题

	Topic string `json:"topic"`
	// **参数说明**：是否加密传输

	EncryptTransport *bool `json:"encrypt_transport,omitempty"`
}

func (o MqsForwarding) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MqsForwarding struct{}"
	}

	return strings.Join([]string{"MqsForwarding", string(data)}, " ")
}
