package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisplayBandwidthPercentageOptions struct {

	// 显示带宽百分比控制量（%）。取值范围为[0-100]。默认：65。
	DisplayBandwidthPercentageValue *int32 `json:"display_bandwidth_percentage_value,omitempty"`
}

func (o DisplayBandwidthPercentageOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisplayBandwidthPercentageOptions struct{}"
	}

	return strings.Join([]string{"DisplayBandwidthPercentageOptions", string(data)}, " ")
}
