package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptions 打印机设备重定向控制的选项。
type PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptions struct {

	// 移动客户端打印机重定向。取值为： false：表示关闭。 true：表示开启。
	MobilePrinterEnable *bool `json:"mobile_printer_enable,omitempty"`

	// 是否开启打印数据压缩模式。取值为： false：表示关闭。 true：表示开启。
	PrintDataCompressionMode *bool `json:"print_data_compression_mode,omitempty"`

	// 是否开启同步客户端默认打印机。取值为： false：表示关闭。 true：表示开启。
	SyncClientDefaultPrinterEnable *bool `json:"sync_client_default_printer_enable,omitempty"`

	// 通用打印机驱动。取值为：Default：linux客户端选择Universal Printing PS，windows客户端选择HDP XPSDrv Driver。HDP XPSDrv Driver。Universal Printing PCL 5。Universal Printing PCL 6。Universal Printing PS。
	UniversalPrinterDriver *PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriver `json:"universal_printer_driver,omitempty"`
}

func (o PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptions struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptions", string(data)}, " ")
}

type PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriver struct {
	value string
}

type PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriverEnum struct {
	DEFAULT                  PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriver
	HDP_XPS_DRV_DRIVER       PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriver
	UNIVERSAL_PRINTING_PCL_5 PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriver
	UNIVERSAL_PRINTING_PCL_6 PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriver
	UNIVERSAL_PRINTING_PS    PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriver
}

func GetPoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriverEnum() PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriverEnum {
	return PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriverEnum{
		DEFAULT: PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriver{
			value: "Default",
		},
		HDP_XPS_DRV_DRIVER: PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriver{
			value: "HDP XPSDrv Driver",
		},
		UNIVERSAL_PRINTING_PCL_5: PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriver{
			value: "Universal Printing PCL 5",
		},
		UNIVERSAL_PRINTING_PCL_6: PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriver{
			value: "Universal Printing PCL 6",
		},
		UNIVERSAL_PRINTING_PS: PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriver{
			value: "Universal Printing PS",
		},
	}
}

func (c PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriver) Value() string {
	return c.value
}

func (c PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriver) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptionsUniversalPrinterDriver) UnmarshalJSON(b []byte) error {
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
