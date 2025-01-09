package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripheralsDeviceRedirectionHidRedirection HID（人体学设备）重定向。
type PoliciesPeripheralsDeviceRedirectionHidRedirection struct {

	// 是否开启HID（人体学设备）重定向。取值为： false：表示关闭。 true：表示开启。
	HidRedirectionEnable *bool `json:"hid_redirection_enable,omitempty"`

	Options *PoliciesPeripheralsDeviceRedirectionHidRedirectionOptions `json:"options,omitempty"`
}

func (o PoliciesPeripheralsDeviceRedirectionHidRedirection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsDeviceRedirectionHidRedirection struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsDeviceRedirectionHidRedirection", string(data)}, " ")
}
