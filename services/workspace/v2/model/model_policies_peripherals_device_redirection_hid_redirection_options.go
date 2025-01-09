package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripheralsDeviceRedirectionHidRedirectionOptions HID控制选项。
type PoliciesPeripheralsDeviceRedirectionHidRedirectionOptions struct {

	// HID重定向自定义策略。
	HidRedirectionCustomizationPolicy *string `json:"hid_redirection_customization_policy,omitempty"`
}

func (o PoliciesPeripheralsDeviceRedirectionHidRedirectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsDeviceRedirectionHidRedirectionOptions struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsDeviceRedirectionHidRedirectionOptions", string(data)}, " ")
}
