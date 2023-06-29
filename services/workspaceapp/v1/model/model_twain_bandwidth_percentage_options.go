package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TwainBandwidthPercentageOptions struct {

	// TWAIN带宽百分比控制量（%）。取值范围为[0-100]。默认：15。
	TwainBandwidthPercentageValue *int32 `json:"twain_bandwidth_percentage_value,omitempty"`
}

func (o TwainBandwidthPercentageOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TwainBandwidthPercentageOptions struct{}"
	}

	return strings.Join([]string{"TwainBandwidthPercentageOptions", string(data)}, " ")
}
