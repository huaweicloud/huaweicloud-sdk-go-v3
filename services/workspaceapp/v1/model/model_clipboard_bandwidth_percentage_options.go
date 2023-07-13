package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClipboardBandwidthPercentageOptions struct {

	// 剪切板带宽百分比控制量（%）。取值范围为[0-100]。默认：3。
	ClipboardBandwidthPercentageValue *int32 `json:"clipboard_bandwidth_percentage_value,omitempty"`
}

func (o ClipboardBandwidthPercentageOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClipboardBandwidthPercentageOptions struct{}"
	}

	return strings.Join([]string{"ClipboardBandwidthPercentageOptions", string(data)}, " ")
}
