package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripheralsUsbPortRedirectionOptions USB端口重定向控制的选项。
type PoliciesPeripheralsUsbPortRedirectionOptions struct {

	// 是否开启图像设备（如: 扫描仪）。取值为： false：表示关闭。 true：表示开启。
	UsbImageEnable *bool `json:"usb_image_enable,omitempty"`

	// 是否开启视频设备（如: 摄像头）。取值为： false：表示关闭。 true：表示开启。
	UsbVideoEnable *bool `json:"usb_video_enable,omitempty"`

	// 是否开启打印设备（如: 打印机）。取值为： false：表示关闭。 true：表示开启。
	UsbPrinterEnable *bool `json:"usb_printer_enable,omitempty"`

	// 是否开启存储设备（如: U盘）。取值为： false：表示关闭。 true：表示开启。
	UsbStorageEnable *bool `json:"usb_storage_enable,omitempty"`

	// 是否开启无线设备（如：蓝牙）。取值为： false：表示关闭。 true：表示开启。
	WirelessDevicesEnable *bool `json:"wireless_devices_enable,omitempty"`

	// 是否开启网路设备（如：无线网卡）。取值为： false：表示关闭。 true：表示开启。
	NetworkDevicesEnable *bool `json:"network_devices_enable,omitempty"`

	// 是否开启智能卡设备（如：Ukey）。取值为： false：表示关闭。 true：表示开启。
	UsbSmartCardEnable *bool `json:"usb_smart_card_enable,omitempty"`

	// 是否开启其他USB设备重定向。取值为： false：表示关闭。 true：表示开启。
	OtherUsbDevicesEnable *bool `json:"other_usb_devices_enable,omitempty"`

	// USB端口重定向自定义策略，长度不能超过18000个字符。
	UsbRedirectionCustomizationPolicy *string `json:"usb_redirection_customization_policy,omitempty"`

	// USB 重定向模式。取值为： 经典模式：Classical mode。 通用模式：Common mode。
	UsbRedirectionMode *string `json:"usb_redirection_mode,omitempty"`
}

func (o PoliciesPeripheralsUsbPortRedirectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsUsbPortRedirectionOptions struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsUsbPortRedirectionOptions", string(data)}, " ")
}
