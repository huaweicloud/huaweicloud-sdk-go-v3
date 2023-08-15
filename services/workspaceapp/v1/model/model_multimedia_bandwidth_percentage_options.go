package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultimediaBandwidthPercentageOptions struct {

	// 多媒体带宽百分比控制量（%）。取值范围为[0-100]。默认：50。
	MultimediaBandwidthPercentageValue *int32 `json:"multimedia_bandwidth_percentage_value,omitempty"`
}

func (o MultimediaBandwidthPercentageOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultimediaBandwidthPercentageOptions struct{}"
	}

	return strings.Join([]string{"MultimediaBandwidthPercentageOptions", string(data)}, " ")
}
