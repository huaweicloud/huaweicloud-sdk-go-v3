package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PcscBandwidthPercentageOptions struct {

	// PCSC带宽百分比控制量（%）。取值范围为[0-100]。默认：5。
	PcscBandwidthPercentageValue *int32 `json:"pcsc_bandwidth_percentage_value,omitempty"`
}

func (o PcscBandwidthPercentageOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PcscBandwidthPercentageOptions struct{}"
	}

	return strings.Join([]string{"PcscBandwidthPercentageOptions", string(data)}, " ")
}
