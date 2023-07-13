package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripheralsUsbDeviceCommonCommonOptions USB和设备智能卡共同控制的选项。
type PoliciesPeripheralsUsbDeviceCommonCommonOptions struct {

	// 是否移除智能卡断开用户会话。取值为： false：表示关闭。 true：表示开启。
	RemoveSmartCardDisconnectEnable *bool `json:"remove_smart_card_disconnect_enable,omitempty"`
}

func (o PoliciesPeripheralsUsbDeviceCommonCommonOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsUsbDeviceCommonCommonOptions struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsUsbDeviceCommonCommonOptions", string(data)}, " ")
}
