package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesDisplayOptionsVideoBitRateOptions 视频码率控制选项。
type PoliciesDisplayOptionsVideoBitRateOptions struct {

	// 视频平均码率（Kbps）。取值范围为[256-100000]。默认：18000。
	AverageVideoBitRate *int32 `json:"average_video_bit_rate,omitempty"`
}

func (o PoliciesDisplayOptionsVideoBitRateOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesDisplayOptionsVideoBitRateOptions struct{}"
	}

	return strings.Join([]string{"PoliciesDisplayOptionsVideoBitRateOptions", string(data)}, " ")
}
