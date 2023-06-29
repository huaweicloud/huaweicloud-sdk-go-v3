package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripheralsSerialPortRedirection 串口重定向。
type PoliciesPeripheralsSerialPortRedirection struct {

	// 是否开启串口重定向。取值为： false：表示关闭。 true：表示开启。
	SerialPortEnable *bool `json:"serial_port_enable,omitempty"`

	Options *PoliciesPeripheralsSerialPortRedirectionOptions `json:"options,omitempty"`
}

func (o PoliciesPeripheralsSerialPortRedirection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsSerialPortRedirection struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsSerialPortRedirection", string(data)}, " ")
}
