package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripherals 外设。
type PoliciesPeripherals struct {

	// 设备调试策略。
	UsbCustomPolicyDebug *string `json:"usb_custom_policy_debug,omitempty"`

	UsbPortRedirection *PoliciesPeripheralsUsbPortRedirection `json:"usb_port_redirection,omitempty"`

	DeviceRedirection *PoliciesPeripheralsDeviceRedirection `json:"device_redirection,omitempty"`

	UsbDeviceCommon *PoliciesPeripheralsUsbDeviceCommon `json:"usb_device_common,omitempty"`

	SerialPortRedirection *PoliciesPeripheralsSerialPortRedirection `json:"serial_port_redirection,omitempty"`

	ParallelPortRedirection *PoliciesPeripheralsParallelPortRedirection `json:"parallel_port_redirection,omitempty"`

	DriverInterfaceRedirection *PoliciesPeripheralsDriverInterfaceRedirection `json:"driver_interface_redirection,omitempty"`
}

func (o PoliciesPeripherals) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripherals struct{}"
	}

	return strings.Join([]string{"PoliciesPeripherals", string(data)}, " ")
}
