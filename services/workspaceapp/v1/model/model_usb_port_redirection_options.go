package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UsbPortRedirectionOptions USB端口重定向控制的选项。
type UsbPortRedirectionOptions struct {

	// 是否开启图像设备（如: 扫描仪）。取值为： - false：表示关闭。 - true：表示开启。
	UsbImageEnable *bool `json:"usb_image_enable,omitempty"`

	// 是否开启视频设备（如: 摄像头）。取值为： - false：表示关闭。 - true：表示开启。
	UsbVideoEnable *bool `json:"usb_video_enable,omitempty"`

	// 是否开启打印设备（如: 打印机）。取值为： - false：表示关闭。 - true：表示开启。
	UsbPrinterEnable *bool `json:"usb_printer_enable,omitempty"`

	// 是否开启存储设备（如: U盘）。取值为： - false：表示关闭。 - true：表示开启。
	UsbStorageEnable *bool `json:"usb_storage_enable,omitempty"`

	// 是否开启无线设备（如：蓝牙）。取值为： - false：表示关闭。 - true：表示开启。
	WirelessDevicesEnable *bool `json:"wireless_devices_enable,omitempty"`

	// 是否开启网络设备（如：无线网卡）。取值为： - false：表示关闭。 - true：表示开启。
	NetworkDevicesEnable *bool `json:"network_devices_enable,omitempty"`

	// 是否开启智能卡设备（如：Ukey）。取值为： - false：表示关闭。 - true：表示开启。
	UsbSmartCardEnable *bool `json:"usb_smart_card_enable,omitempty"`

	// 是否开启其他USB设备重定向。取值为： - false：表示关闭。 - true：表示开启。
	OtherUsbDevicesEnable *bool `json:"other_usb_devices_enable,omitempty"`

	// USB端口重定向自定义策略。
	UsbRedirectionCustomizationPolicy *string `json:"usb_redirection_customization_policy,omitempty"`

	// USB 重定向模式。取值为： - Classical mode：经典模式。 - Common mode：通用模式。
	UsbRedirectionMode *UsbPortRedirectionOptionsUsbRedirectionMode `json:"usb_redirection_mode,omitempty"`
}

func (o UsbPortRedirectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UsbPortRedirectionOptions struct{}"
	}

	return strings.Join([]string{"UsbPortRedirectionOptions", string(data)}, " ")
}

type UsbPortRedirectionOptionsUsbRedirectionMode struct {
	value string
}

type UsbPortRedirectionOptionsUsbRedirectionModeEnum struct {
	CLASSICAL_MODE UsbPortRedirectionOptionsUsbRedirectionMode
	COMMON_MODE    UsbPortRedirectionOptionsUsbRedirectionMode
}

func GetUsbPortRedirectionOptionsUsbRedirectionModeEnum() UsbPortRedirectionOptionsUsbRedirectionModeEnum {
	return UsbPortRedirectionOptionsUsbRedirectionModeEnum{
		CLASSICAL_MODE: UsbPortRedirectionOptionsUsbRedirectionMode{
			value: "Classical mode",
		},
		COMMON_MODE: UsbPortRedirectionOptionsUsbRedirectionMode{
			value: "Common mode",
		},
	}
}

func (c UsbPortRedirectionOptionsUsbRedirectionMode) Value() string {
	return c.value
}

func (c UsbPortRedirectionOptionsUsbRedirectionMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UsbPortRedirectionOptionsUsbRedirectionMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
