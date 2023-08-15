package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrinterBandwidthPercentageOptions struct {

	// 打印机带宽百分比控制量（%）。取值范围为[0-100]。默认：5。
	PrinterBandwidthPercentageValue *int32 `json:"printer_bandwidth_percentage_value,omitempty"`
}

func (o PrinterBandwidthPercentageOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrinterBandwidthPercentageOptions struct{}"
	}

	return strings.Join([]string{"PrinterBandwidthPercentageOptions", string(data)}, " ")
}
