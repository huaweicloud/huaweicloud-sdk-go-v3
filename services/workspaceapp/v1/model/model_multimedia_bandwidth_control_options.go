package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultimediaBandwidthControlOptions struct {

	// 多媒体带宽控制量（Kbps）。取值范围为[5000-20000]。默认：15000。
	MultimediaBandwidthControlValue *int32 `json:"multimedia_bandwidth_control_value,omitempty"`
}

func (o MultimediaBandwidthControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultimediaBandwidthControlOptions struct{}"
	}

	return strings.Join([]string{"MultimediaBandwidthControlOptions", string(data)}, " ")
}
