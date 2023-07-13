package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimeoutConfig 超时配置，开启后不支持关闭
type TimeoutConfig struct {

	// WAF连接源站超时配置
	ConnectTimeout *int32 `json:"connect_timeout,omitempty"`

	// WAF发送请求到源站超时配置
	SendTimeout *int32 `json:"send_timeout,omitempty"`

	// WAF接收源站响应超时配置
	ReadTimeout *int32 `json:"read_timeout,omitempty"`
}

func (o TimeoutConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeoutConfig struct{}"
	}

	return strings.Join([]string{"TimeoutConfig", string(data)}, " ")
}
