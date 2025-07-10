package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
	UniversalPrinterDriver *string `json:"universal_printer_driver,omitempty"`
}

func (o PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptions struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsDeviceRedirectionPrinterRedirectionOptions", string(data)}, " ")
}
