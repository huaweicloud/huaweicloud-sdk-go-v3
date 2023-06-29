package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UsbBandwidthControlOptions struct {

	// USB带宽控制量（Kbps）。取值范围为[1000-30000]。默认：30000。
	UsbBandwidthControlValue *int32 `json:"usb_bandwidth_control_value,omitempty"`
}

func (o UsbBandwidthControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UsbBandwidthControlOptions struct{}"
	}

	return strings.Join([]string{"UsbBandwidthControlOptions", string(data)}, " ")
}
