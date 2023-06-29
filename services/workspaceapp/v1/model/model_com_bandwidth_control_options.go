package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComBandwidthControlOptions struct {

	// 串口带宽控制量（Kbps）。取值范围为[500-2000]。默认：1000。
	ComBandwidthControlValue *int32 `json:"com_bandwidth_control_value,omitempty"`
}

func (o ComBandwidthControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComBandwidthControlOptions struct{}"
	}

	return strings.Join([]string{"ComBandwidthControlOptions", string(data)}, " ")
}
