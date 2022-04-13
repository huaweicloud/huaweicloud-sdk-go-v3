package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 频率模式配置。和periods字段二选一，不可共存。
type TaskTimingFrequency struct {
	// 相邻两次执行之间的间隔，频率模式必填。取值范围在5~720之间，单位：分钟。

	Interval int32 `json:"interval"`
	// 单次执行的运行时长，频率模式必填。取值范围在5~720之间，单位：分钟。

	Duration int32 `json:"duration"`
	// 单日内执行的起始时间，选填。格式形如hh:mm:ss。

	BeginAt *string `json:"begin_at,omitempty"`
	// 单日内执行的结束时间，选填。格式形如hh:mm:ss。

	EndAt *string `json:"end_at,omitempty"`
}

func (o TaskTimingFrequency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskTimingFrequency struct{}"
	}

	return strings.Join([]string{"TaskTimingFrequency", string(data)}, " ")
}
