package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripheralsUsbDeviceCommon USB和设备智能卡控制选项。
type PoliciesPeripheralsUsbDeviceCommon struct {

	// 是否开启PC/SC智能卡重定向。取值为： Enable：表示已启动。 Closed：表示已关闭。 Disable：表示已禁用。
	PcscSmartCardEnable *string `json:"pcsc_smart_card_enable,omitempty"`

	CommonOptions *PoliciesPeripheralsUsbDeviceCommonCommonOptions `json:"common_options,omitempty"`
}

func (o PoliciesPeripheralsUsbDeviceCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsUsbDeviceCommon struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsUsbDeviceCommon", string(data)}, " ")
}
