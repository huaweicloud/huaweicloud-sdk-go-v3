package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SessionPrinterOptions 会话打印机控制选项。
type SessionPrinterOptions struct {

	// 会话打印机自定义策略。
	SessionPrinterCustomizationPolicy *string `json:"session_printer_customization_policy,omitempty"`
}

func (o SessionPrinterOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionPrinterOptions struct{}"
	}

	return strings.Join([]string{"SessionPrinterOptions", string(data)}, " ")
}
