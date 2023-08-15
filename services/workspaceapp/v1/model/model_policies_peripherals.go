package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripherals 外设。
type PoliciesPeripherals struct {
	UsbPortRedirection *PoliciesPeripheralsUsbPortRedirection `json:"usb_port_redirection,omitempty"`

	DeviceRedirection *PoliciesPeripheralsDeviceRedirection `json:"device_redirection,omitempty"`

	UsbDeviceCommon *PoliciesPeripheralsUsbDeviceCommon `json:"usb_device_common,omitempty"`

	SerialPortRedirection *PoliciesPeripheralsSerialPortRedirection `json:"serial_port_redirection,omitempty"`
}

func (o PoliciesPeripherals) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripherals struct{}"
	}

	return strings.Join([]string{"PoliciesPeripherals", string(data)}, " ")
}
