package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripheralsDeviceRedirectionSessionPrinter 会话打印机。
type PoliciesPeripheralsDeviceRedirectionSessionPrinter struct {

	// 是否开启会话打印机。取值为： false：表示关闭。 true：表示开启。
	SessionPrinterEnable *bool `json:"session_printer_enable,omitempty"`

	Options *PoliciesPeripheralsDeviceRedirectionSessionPrinterOptions `json:"options,omitempty"`
}

func (o PoliciesPeripheralsDeviceRedirectionSessionPrinter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsDeviceRedirectionSessionPrinter struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsDeviceRedirectionSessionPrinter", string(data)}, " ")
}
