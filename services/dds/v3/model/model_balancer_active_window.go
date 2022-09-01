package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 活动时间窗设置。
type BalancerActiveWindow struct {

	// 活动时间窗开始时间。
	StartTime *string `json:"start_time,omitempty" xml:"start_time"`

	// 活动时间窗结束时间。
	StopTime *string `json:"stop_time,omitempty" xml:"stop_time"`
}

func (o BalancerActiveWindow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BalancerActiveWindow struct{}"
	}

	return strings.Join([]string{"BalancerActiveWindow", string(data)}, " ")
}
