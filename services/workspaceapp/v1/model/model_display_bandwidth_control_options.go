package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisplayBandwidthControlOptions struct {

	// 显示带宽控制量（Kbps）。取值范围为[500-50000]。默认：20000。
	DisplayBandwidthControlValue *int32 `json:"display_bandwidth_control_value,omitempty"`
}

func (o DisplayBandwidthControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisplayBandwidthControlOptions struct{}"
	}

	return strings.Join([]string{"DisplayBandwidthControlOptions", string(data)}, " ")
}
