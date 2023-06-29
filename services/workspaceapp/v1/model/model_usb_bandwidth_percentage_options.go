package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UsbBandwidthPercentageOptions struct {

	// USB带宽百分比控制量（%）。取值范围为[0-100]。默认：100。
	UsbBandwidthPercentageValue *int32 `json:"usb_bandwidth_percentage_value,omitempty"`
}

func (o UsbBandwidthPercentageOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UsbBandwidthPercentageOptions struct{}"
	}

	return strings.Join([]string{"UsbBandwidthPercentageOptions", string(data)}, " ")
}
