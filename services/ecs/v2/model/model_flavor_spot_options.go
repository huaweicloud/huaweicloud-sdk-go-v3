package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorSpotOptions
type FlavorSpotOptions struct {

	// 购买的竞价实例时长
	LongestSpotDurationHours *int32 `json:"longest_spot_duration_hours,omitempty"`

	// 购买的“竞价实例时长”的个数。
	LargestSpotDurationCount *int32 `json:"largest_spot_duration_count,omitempty"`

	// 竞价实例中断策略。  - immediate：立即释放 - delay：延迟释放
	InterruptionPolicy *string `json:"interruption_policy,omitempty"`
}

func (o FlavorSpotOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorSpotOptions struct{}"
	}

	return strings.Join([]string{"FlavorSpotOptions", string(data)}, " ")
}
