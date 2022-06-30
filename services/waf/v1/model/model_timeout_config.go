package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 超时配置
type TimeoutConfig struct {

	// 连接超时配置(秒)
	ConnectTimeout *int32 `json:"connect_timeout,omitempty"`

	// 发送超时配置(秒)
	SendTimeout *int32 `json:"send_timeout,omitempty"`

	// 接收超时配置(秒)
	ReadTimeout *int32 `json:"read_timeout,omitempty"`
}

func (o TimeoutConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeoutConfig struct{}"
	}

	return strings.Join([]string{"TimeoutConfig", string(data)}, " ")
}
