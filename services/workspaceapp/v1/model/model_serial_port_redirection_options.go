package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SerialPortRedirectionOptions 串口重定向控制的选项。
type SerialPortRedirectionOptions struct {

	// 是否自动连接客户端串口。取值为： false：表示关闭。 true：表示开启。
	AutoConnectEnable *bool `json:"auto_connect_enable,omitempty"`
}

func (o SerialPortRedirectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SerialPortRedirectionOptions struct{}"
	}

	return strings.Join([]string{"SerialPortRedirectionOptions", string(data)}, " ")
}
