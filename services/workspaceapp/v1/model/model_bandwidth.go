package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Bandwidth struct {

	// 智能显示传输。取值为：DISABLE：表示关闭。ENABLE：表示开启。DIAGNOSTIC：诊断模式。
	IntelligentDataTransportFlag *BandwidthIntelligentDataTransportFlag `json:"intelligent_data_transport_flag,omitempty"`

	// 是否开启总带宽控制。取值为：false：表示关闭。true：表示开启。
	TotalBandwidthControlEnable *bool `json:"total_bandwidth_control_enable,omitempty"`

	Options *TotalBandwidthControlOptions `json:"options,omitempty"`

	// 是否开启显示带宽控制。取值为：false：表示关闭。true：表示开启。
	DisplayBandwidthControlEnable *bool `json:"display_bandwidth_control_enable,omitempty"`

	DisplayBandwidthControlOptions *DisplayBandwidthControlOptions `json:"display_bandwidth_control_options,omitempty"`

	// 是否开启多媒体带宽控制。取值为：false：表示关闭。true：表示开启。
	MultimediaBandwidthControlEnable *bool `json:"multimedia_bandwidth_control_enable,omitempty"`

	MultimediaBandwidthControlOptions *MultimediaBandwidthControlOptions `json:"multimedia_bandwidth_control_options,omitempty"`

	// 是否开启USB带宽控制。取值为：false：表示关闭。true：表示开启。
	UsbBandwidthControlEnable *bool `json:"usb_bandwidth_control_enable,omitempty"`

	UsbBandwidthControlOptions *UsbBandwidthControlOptions `json:"usb_bandwidth_control_options,omitempty"`

	// 是否开启PCSC控制。取值为：false：表示关闭。true：表示开启。
	PcscBandwidthControlEnable *bool `json:"pcsc_bandwidth_control_enable,omitempty"`

	PcscBandwidthControlOptions *PcscBandwidthControlOptions `json:"pcsc_bandwidth_control_options,omitempty"`

	// 是否开启TWAIN带宽控制。取值为：false：表示关闭。true：表示开启。
	TwainBandwidthControlEnable *bool `json:"twain_bandwidth_control_enable,omitempty"`

	TwainBandwidthControlOptions *TwainBandwidthControlOptions `json:"twain_bandwidth_control_options,omitempty"`

	// 是否开启打印机带宽控制。取值为：false：表示关闭。true：表示开启。
	PrinterBandwidthControlEnable *bool `json:"printer_bandwidth_control_enable,omitempty"`

	PrinterBandwidthControlOptions *PrinterBandwidthControlOptions `json:"printer_bandwidth_control_options,omitempty"`

	// 是否开启串口带宽控制。取值为：false：表示关闭。true：表示开启。
	ComBandwidthControlEnable *bool `json:"com_bandwidth_control_enable,omitempty"`

	ComBandwidthControlOptions *ComBandwidthControlOptions `json:"com_bandwidth_control_options,omitempty"`

	// 是否开启文件重定向带宽控制。取值为：false：表示关闭。true：表示开启。
	FileRedirectionBandwidthControlEnable *bool `json:"file_redirection_bandwidth_control_enable,omitempty"`

	FileRedirectionBandwidthControlOptions *FileRedirectionBandwidthControlOptions `json:"file_redirection_bandwidth_control_options,omitempty"`

	// 是否开启剪切板带宽控制。取值为：false：表示关闭。true：表示开启。
	ClipboardBandwidthControlEnable *bool `json:"clipboard_bandwidth_control_enable,omitempty"`

	ClipboardBandwidthControlOptions *ClipboardBandwidthControlOptions `json:"clipboard_bandwidth_control_options,omitempty"`

	// 是否开启安全通道带宽控制。取值为：false：表示关闭。true：表示开启。
	SecureChannelBandwidthControlEnable *bool `json:"secure_channel_bandwidth_control_enable,omitempty"`

	SecureChannelBandwidthControlOptions *SecureChannelBandwidthControlOptions `json:"secure_channel_bandwidth_control_options,omitempty"`

	// 是否开启摄像头带宽控制。取值为：false：表示关闭。true：表示开启。
	CameraBandwidthControlEnable *bool `json:"camera_bandwidth_control_enable,omitempty"`

	CameraBandwidthControlOptions *CameraBandwidthControlOptions `json:"camera_bandwidth_control_options,omitempty"`

	// 是否开启虚拟通道带宽控制。取值为：false：表示关闭。true：表示开启。
	VirtualChannelBandwidthControlEnable *bool `json:"virtual_channel_bandwidth_control_enable,omitempty"`

	VirtualChannelBandwidthControlOptions *VirtualChannelBandwidthControlOptions `json:"virtual_channel_bandwidth_control_options,omitempty"`
}

func (o Bandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Bandwidth struct{}"
	}

	return strings.Join([]string{"Bandwidth", string(data)}, " ")
}

type BandwidthIntelligentDataTransportFlag struct {
	value string
}

type BandwidthIntelligentDataTransportFlagEnum struct {
	DISABLE    BandwidthIntelligentDataTransportFlag
	ENABLE     BandwidthIntelligentDataTransportFlag
	DIAGNOSTIC BandwidthIntelligentDataTransportFlag
}

func GetBandwidthIntelligentDataTransportFlagEnum() BandwidthIntelligentDataTransportFlagEnum {
	return BandwidthIntelligentDataTransportFlagEnum{
		DISABLE: BandwidthIntelligentDataTransportFlag{
			value: "DISABLE",
		},
		ENABLE: BandwidthIntelligentDataTransportFlag{
			value: "ENABLE",
		},
		DIAGNOSTIC: BandwidthIntelligentDataTransportFlag{
			value: "DIAGNOSTIC",
		},
	}
}

func (c BandwidthIntelligentDataTransportFlag) Value() string {
	return c.value
}

func (c BandwidthIntelligentDataTransportFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BandwidthIntelligentDataTransportFlag) UnmarshalJSON(b []byte) error {
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
