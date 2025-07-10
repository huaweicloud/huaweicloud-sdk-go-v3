package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
	ImageCompressionLevel *string `json:"image_compression_level,omitempty"`

	HidRedirection *PoliciesPeripheralsDeviceRedirectionHidRedirection `json:"hid_redirection,omitempty"`
}

func (o PoliciesPeripheralsDeviceRedirection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsDeviceRedirection struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsDeviceRedirection", string(data)}, " ")
}
