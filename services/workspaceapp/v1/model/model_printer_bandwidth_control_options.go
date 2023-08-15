package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrinterBandwidthControlOptions struct {

	// 打印机带宽控制量（Kbps）。取值范围为[1000-5000]。默认：2000。
	PrinterBandwidthControlValue *int32 `json:"printer_bandwidth_control_value,omitempty"`
}

func (o PrinterBandwidthControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrinterBandwidthControlOptions struct{}"
	}

	return strings.Join([]string{"PrinterBandwidthControlOptions", string(data)}, " ")
}
