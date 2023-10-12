package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PrinterRedirectionOptions 打印机设备重定向控制的选项。
type PrinterRedirectionOptions struct {

	// 是否开启同步客户端默认打印机。取值为： false：表示关闭。 true：表示开启。
	SyncClientDefaultPrinterEnable *bool `json:"sync_client_default_printer_enable,omitempty"`

	// 通用打印机驱动。取值为：- Default：linux客户端选择Universal Printing- PS，windows客户端选择HDP XPSDrv Driver。- HDP XPSDrv Driver。- Universal Printing PCL 5。- Universal Printing PCL 6。- Universal Printing PS。
	UniversalPrinterDriver *PrinterRedirectionOptionsUniversalPrinterDriver `json:"universal_printer_driver,omitempty"`
}

func (o PrinterRedirectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrinterRedirectionOptions struct{}"
	}

	return strings.Join([]string{"PrinterRedirectionOptions", string(data)}, " ")
}

type PrinterRedirectionOptionsUniversalPrinterDriver struct {
	value string
}

type PrinterRedirectionOptionsUniversalPrinterDriverEnum struct {
	DEFAULT                  PrinterRedirectionOptionsUniversalPrinterDriver
	HDP_XPS_DRV_DRIVER       PrinterRedirectionOptionsUniversalPrinterDriver
	UNIVERSAL_PRINTING_PCL_5 PrinterRedirectionOptionsUniversalPrinterDriver
	UNIVERSAL_PRINTING_PCL_6 PrinterRedirectionOptionsUniversalPrinterDriver
	UNIVERSAL_PRINTING_PS    PrinterRedirectionOptionsUniversalPrinterDriver
}

func GetPrinterRedirectionOptionsUniversalPrinterDriverEnum() PrinterRedirectionOptionsUniversalPrinterDriverEnum {
	return PrinterRedirectionOptionsUniversalPrinterDriverEnum{
		DEFAULT: PrinterRedirectionOptionsUniversalPrinterDriver{
			value: "Default",
		},
		HDP_XPS_DRV_DRIVER: PrinterRedirectionOptionsUniversalPrinterDriver{
			value: "HDP XPSDrv Driver",
		},
		UNIVERSAL_PRINTING_PCL_5: PrinterRedirectionOptionsUniversalPrinterDriver{
			value: "Universal Printing PCL 5",
		},
		UNIVERSAL_PRINTING_PCL_6: PrinterRedirectionOptionsUniversalPrinterDriver{
			value: "Universal Printing PCL 6",
		},
		UNIVERSAL_PRINTING_PS: PrinterRedirectionOptionsUniversalPrinterDriver{
			value: "Universal Printing PS",
		},
	}
}

func (c PrinterRedirectionOptionsUniversalPrinterDriver) Value() string {
	return c.value
}

func (c PrinterRedirectionOptionsUniversalPrinterDriver) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrinterRedirectionOptionsUniversalPrinterDriver) UnmarshalJSON(b []byte) error {
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
