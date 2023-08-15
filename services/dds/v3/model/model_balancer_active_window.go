package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BalancerActiveWindow 活动时间窗设置。
type BalancerActiveWindow struct {

	// 活动时间窗开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 活动时间窗结束时间。
	StopTime *string `json:"stop_time,omitempty"`
}

func (o BalancerActiveWindow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BalancerActiveWindow struct{}"
	}

	return strings.Join([]string{"BalancerActiveWindow", string(data)}, " ")
}
