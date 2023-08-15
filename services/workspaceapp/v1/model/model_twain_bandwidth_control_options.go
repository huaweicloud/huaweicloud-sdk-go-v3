package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TwainBandwidthControlOptions struct {

	// TWAIN带宽控制量（Kbps）。取值范围为[2000-10000]。默认：5000。
	TwainBandwidthControlValue *int32 `json:"twain_bandwidth_control_value,omitempty"`
}

func (o TwainBandwidthControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TwainBandwidthControlOptions struct{}"
	}

	return strings.Join([]string{"TwainBandwidthControlOptions", string(data)}, " ")
}
