package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PoliciesPeripheralsUsbDeviceCommon USB和设备智能卡控制选项。
type PoliciesPeripheralsUsbDeviceCommon struct {

	// 是否开启PC/SC智能卡重定向。取值为： Enable：表示已启动。 Closed：表示已关闭。 Disable：表示已禁用。
	PcscSmartCardEnable *PoliciesPeripheralsUsbDeviceCommonPcscSmartCardEnable `json:"pcsc_smart_card_enable,omitempty"`

	CommonOptions *PoliciesPeripheralsUsbDeviceCommonCommonOptions `json:"common_options,omitempty"`
}

func (o PoliciesPeripheralsUsbDeviceCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsUsbDeviceCommon struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsUsbDeviceCommon", string(data)}, " ")
}

type PoliciesPeripheralsUsbDeviceCommonPcscSmartCardEnable struct {
	value string
}

type PoliciesPeripheralsUsbDeviceCommonPcscSmartCardEnableEnum struct {
	ENABLE  PoliciesPeripheralsUsbDeviceCommonPcscSmartCardEnable
	CLOSED  PoliciesPeripheralsUsbDeviceCommonPcscSmartCardEnable
	DISABLE PoliciesPeripheralsUsbDeviceCommonPcscSmartCardEnable
}

func GetPoliciesPeripheralsUsbDeviceCommonPcscSmartCardEnableEnum() PoliciesPeripheralsUsbDeviceCommonPcscSmartCardEnableEnum {
	return PoliciesPeripheralsUsbDeviceCommonPcscSmartCardEnableEnum{
		ENABLE: PoliciesPeripheralsUsbDeviceCommonPcscSmartCardEnable{
			value: "Enable",
		},
		CLOSED: PoliciesPeripheralsUsbDeviceCommonPcscSmartCardEnable{
			value: "Closed",
		},
		DISABLE: PoliciesPeripheralsUsbDeviceCommonPcscSmartCardEnable{
			value: "Disable",
		},
	}
}

func (c PoliciesPeripheralsUsbDeviceCommonPcscSmartCardEnable) Value() string {
	return c.value
}

func (c PoliciesPeripheralsUsbDeviceCommonPcscSmartCardEnable) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesPeripheralsUsbDeviceCommonPcscSmartCardEnable) UnmarshalJSON(b []byte) error {
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
