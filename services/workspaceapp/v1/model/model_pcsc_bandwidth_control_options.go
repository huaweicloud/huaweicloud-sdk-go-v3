package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PcscBandwidthControlOptions struct {

	// PCSC带宽控制量（Kbps）。取值范围为[1000-5000]。默认：2000。
	PcscBandwidthControlValue *int32 `json:"pcsc_bandwidth_control_value,omitempty"`
}

func (o PcscBandwidthControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PcscBandwidthControlOptions struct{}"
	}

	return strings.Join([]string{"PcscBandwidthControlOptions", string(data)}, " ")
}
