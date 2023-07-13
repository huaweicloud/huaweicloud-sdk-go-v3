package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripheralsDeviceRedirectionPrinterRedirection 打印机重定向。
type PoliciesPeripheralsDeviceRedirectionPrinterRedirection struct {

	// 是否开启打印机设备重定向。取值为： - false：表示关闭。 - true：表示开启。
	PrinterEnable *bool `json:"printer_enable,omitempty"`

	Options *PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptions `json:"options,omitempty"`
}

func (o PoliciesPeripheralsDeviceRedirectionPrinterRedirection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsDeviceRedirectionPrinterRedirection struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsDeviceRedirectionPrinterRedirection", string(data)}, " ")
}
