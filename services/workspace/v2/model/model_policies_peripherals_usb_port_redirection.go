package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripheralsUsbPortRedirection USB端口重定向。
type PoliciesPeripheralsUsbPortRedirection struct {

	// 是否开启USB端口重定向。 false：表示关闭。 true：表示开启。
	UsbEnable *bool `json:"usb_enable,omitempty"`

	Options *PoliciesPeripheralsUsbPortRedirectionOptions `json:"options,omitempty"`
}

func (o PoliciesPeripheralsUsbPortRedirection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsUsbPortRedirection struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsUsbPortRedirection", string(data)}, " ")
}
