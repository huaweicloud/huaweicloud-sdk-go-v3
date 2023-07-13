package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripheralsSerialPortRedirectionOptions 串口重定向控制的选项。
type PoliciesPeripheralsSerialPortRedirectionOptions struct {

	// 是否自动连接客户端串口。取值为： false：表示关闭。 true：表示开启。
	AutoConnectEnable *bool `json:"auto_connect_enable,omitempty"`
}

func (o PoliciesPeripheralsSerialPortRedirectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsSerialPortRedirectionOptions struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsSerialPortRedirectionOptions", string(data)}, " ")
}
