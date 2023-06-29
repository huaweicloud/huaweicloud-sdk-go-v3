package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComBandwidthPercentageOptions struct {

	// 串口带宽百分比控制量（%）。取值范围为[0-100]。默认：3。
	ComBandwidthPercentageValue *int32 `json:"com_bandwidth_percentage_value,omitempty"`
}

func (o ComBandwidthPercentageOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComBandwidthPercentageOptions struct{}"
	}

	return strings.Join([]string{"ComBandwidthPercentageOptions", string(data)}, " ")
}
