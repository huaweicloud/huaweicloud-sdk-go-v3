package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripheralsDeviceRedirectionSessionPrinterOptions 会话打印机控制选项。
type PoliciesPeripheralsDeviceRedirectionSessionPrinterOptions struct {

	// 会话打印机自定义策略。
	SessionPrinterCustomizationPolicy *string `json:"session_printer_customization_policy,omitempty"`
}

func (o PoliciesPeripheralsDeviceRedirectionSessionPrinterOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsDeviceRedirectionSessionPrinterOptions struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsDeviceRedirectionSessionPrinterOptions", string(data)}, " ")
}
