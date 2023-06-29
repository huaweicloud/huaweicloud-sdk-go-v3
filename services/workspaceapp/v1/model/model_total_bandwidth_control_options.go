package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TotalBandwidthControlOptions struct {

	// 总带宽控制量（Kbps）。取值范围为[10000-1000000]。默认：30000。
	TotalBandwidthControlValue *int32 `json:"total_bandwidth_control_value,omitempty"`

	// 显示带宽百分比控制。取值为：false：表示关闭。true：表示开启。
	DisplayBandwidthPercentageEnable *bool `json:"display_bandwidth_percentage_enable,omitempty"`

	DisplayBandwidthPercentageOptions *DisplayBandwidthPercentageOptions `json:"display_bandwidth_percentage_options,omitempty"`

	// 多媒体带宽百分比控制。取值为：false：表示关闭。true：表示开启。
	MultimediaBandwidthPercentageEnable *bool `json:"multimedia_bandwidth_percentage_enable,omitempty"`

	MultimediaBandwidthPercentageOptions *MultimediaBandwidthPercentageOptions `json:"multimedia_bandwidth_percentage_options,omitempty"`

	// USB带宽百分比控制。取值为：false：表示关闭。true：表示开启。
	UsbBandwidthPercentageEnable *bool `json:"usb_bandwidth_percentage_enable,omitempty"`

	UsbBandwidthPercentageOptions *UsbBandwidthPercentageOptions `json:"usb_bandwidth_percentage_options,omitempty"`

	// PCSC带宽百分比控制。取值为：false：表示关闭。true：表示开启。
	PcscBandwidthPercentageEnable *bool `json:"pcsc_bandwidth_percentage_enable,omitempty"`

	PcscBandwidthPercentageOptions *PcscBandwidthPercentageOptions `json:"pcsc_bandwidth_percentage_options,omitempty"`

	// TWAIN带宽百分比控制。取值为：false：表示关闭。true：表示开启。
	TwainBandwidthPercentageEnable *bool `json:"twain_bandwidth_percentage_enable,omitempty"`

	TwainBandwidthPercentageOptions *TwainBandwidthPercentageOptions `json:"twain_bandwidth_percentage_options,omitempty"`

	// 打印机带宽百分比控制。取值为：false：表示关闭。true：表示开启。
	PrinterBandwidthPercentageEnable *bool `json:"printer_bandwidth_percentage_enable,omitempty"`

	PrinterBandwidthPercentageOptions *PrinterBandwidthPercentageOptions `json:"printer_bandwidth_percentage_options,omitempty"`

	// 串口带宽百分比控制。取值为：false：表示关闭。true：表示开启。
	ComBandwidthPercentageEnable *bool `json:"com_bandwidth_percentage_enable,omitempty"`

	ComBandwidthPercentageOptions *ComBandwidthPercentageOptions `json:"com_bandwidth_percentage_options,omitempty"`

	// 文件重定向带宽百分比控制。取值为：false：表示关闭。true：表示开启。
	FileRedirectionBandwidthPercentageEnable *bool `json:"file_redirection_bandwidth_percentage_enable,omitempty"`

	FileRedirectionBandwidthPercentageOptions *FileRedirectionBandwidthPercentageOptions `json:"file_redirection_bandwidth_percentage_options,omitempty"`

	// 剪切板带宽百分比控制。取值为：false：表示关闭。true：表示开启。
	ClipboardBandwidthPercentageEnable *bool `json:"clipboard_bandwidth_percentage_enable,omitempty"`

	ClipboardBandwidthPercentageOptions *ClipboardBandwidthPercentageOptions `json:"clipboard_bandwidth_percentage_options,omitempty"`

	// 安全通道带宽百分比控制。取值为：false：表示关闭。true：表示开启。
	SecureChannelBandwidthPercentageEnable *bool `json:"secure_channel_bandwidth_percentage_enable,omitempty"`

	SecureChannelBandwidthPercentageOptions *SecureChannelBandwidthPercentageOptions `json:"secure_channel_bandwidth_percentage_options,omitempty"`

	// 摄像头带宽百分比控制。取值为：false：表示关闭。true：表示开启。
	CameraBandwidthPercentageEnable *bool `json:"camera_bandwidth_percentage_enable,omitempty"`

	CameraBandwidthPercentageOptions *CameraBandwidthPercentageOptions `json:"camera_bandwidth_percentage_options,omitempty"`

	// 虚拟通道带宽百分比控制。取值为：false：表示关闭。true：表示开启。
	VirtualChannelBandwidthPercentageEnable *bool `json:"virtual_channel_bandwidth_percentage_enable,omitempty"`

	VirtualChannelBandwidthPercentageOptions *VirtualChannelBandwidthPercentageOptions `json:"virtual_channel_bandwidth_percentage_options,omitempty"`
}

func (o TotalBandwidthControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TotalBandwidthControlOptions struct{}"
	}

	return strings.Join([]string{"TotalBandwidthControlOptions", string(data)}, " ")
}
