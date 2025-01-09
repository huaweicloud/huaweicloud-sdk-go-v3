package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PoliciesPeripheralsDeviceRedirection 设备重定向。
type PoliciesPeripheralsDeviceRedirection struct {
	PrinterRedirection *PoliciesPeripheralsDeviceRedirectionPrinterRedirection `json:"printer_redirection,omitempty"`

	SessionPrinter *PoliciesPeripheralsDeviceRedirectionSessionPrinter `json:"session_printer,omitempty"`

	CameraRedirection *PoliciesPeripheralsDeviceRedirectionCameraRedirection `json:"camera_redirection,omitempty"`

	// 是否开启TWAIN设备重定向。取值为： false：表示关闭。 true：表示开启。
	TwainRedirectionEnable *bool `json:"twain_redirection_enable,omitempty"`

	// 图形压缩级别。取值为： 不压缩：none。 低（速度最快）：low。 中（速度适中）：medium。 高（速度最慢）：high。 无损（无损压缩）：lossless。 低损（低损压缩）：low-loss。 中损（中损压缩）：medium-loss。 高损（高损压缩）：high-loss。
	ImageCompressionLevel *PoliciesPeripheralsDeviceRedirectionImageCompressionLevel `json:"image_compression_level,omitempty"`

	HidRedirection *PoliciesPeripheralsDeviceRedirectionHidRedirection `json:"hid_redirection,omitempty"`
}

func (o PoliciesPeripheralsDeviceRedirection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsDeviceRedirection struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsDeviceRedirection", string(data)}, " ")
}

type PoliciesPeripheralsDeviceRedirectionImageCompressionLevel struct {
	value string
}

type PoliciesPeripheralsDeviceRedirectionImageCompressionLevelEnum struct {
	NONE        PoliciesPeripheralsDeviceRedirectionImageCompressionLevel
	LOW         PoliciesPeripheralsDeviceRedirectionImageCompressionLevel
	MEDIUM      PoliciesPeripheralsDeviceRedirectionImageCompressionLevel
	HIGH        PoliciesPeripheralsDeviceRedirectionImageCompressionLevel
	LOSSLESS    PoliciesPeripheralsDeviceRedirectionImageCompressionLevel
	LOW_LOSS    PoliciesPeripheralsDeviceRedirectionImageCompressionLevel
	MEDIUM_LOSS PoliciesPeripheralsDeviceRedirectionImageCompressionLevel
	HIGH_LOSS   PoliciesPeripheralsDeviceRedirectionImageCompressionLevel
}

func GetPoliciesPeripheralsDeviceRedirectionImageCompressionLevelEnum() PoliciesPeripheralsDeviceRedirectionImageCompressionLevelEnum {
	return PoliciesPeripheralsDeviceRedirectionImageCompressionLevelEnum{
		NONE: PoliciesPeripheralsDeviceRedirectionImageCompressionLevel{
			value: "none",
		},
		LOW: PoliciesPeripheralsDeviceRedirectionImageCompressionLevel{
			value: "low",
		},
		MEDIUM: PoliciesPeripheralsDeviceRedirectionImageCompressionLevel{
			value: "medium",
		},
		HIGH: PoliciesPeripheralsDeviceRedirectionImageCompressionLevel{
			value: "high",
		},
		LOSSLESS: PoliciesPeripheralsDeviceRedirectionImageCompressionLevel{
			value: "lossless",
		},
		LOW_LOSS: PoliciesPeripheralsDeviceRedirectionImageCompressionLevel{
			value: "low-loss",
		},
		MEDIUM_LOSS: PoliciesPeripheralsDeviceRedirectionImageCompressionLevel{
			value: "medium-loss",
		},
		HIGH_LOSS: PoliciesPeripheralsDeviceRedirectionImageCompressionLevel{
			value: "high-loss",
		},
	}
}

func (c PoliciesPeripheralsDeviceRedirectionImageCompressionLevel) Value() string {
	return c.value
}

func (c PoliciesPeripheralsDeviceRedirectionImageCompressionLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesPeripheralsDeviceRedirectionImageCompressionLevel) UnmarshalJSON(b []byte) error {
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
