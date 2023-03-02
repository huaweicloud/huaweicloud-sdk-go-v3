package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 事件订阅监控指标数据
type EventSubMetricsItem struct {

	// 失败数
	FailNum *int64 `json:"fail_num,omitempty"`

	// 重试成功数
	RetrySuccessNum *int64 `json:"retry_success_num,omitempty"`

	// 重试失败数
	RetryFailNum *int64 `json:"retry_fail_num,omitempty"`

	// 重试数
	RetryTimes *int32 `json:"retry_times,omitempty"`

	// 时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 调用数
	Num *int64 `json:"num,omitempty"`

	// 调用成功数
	SuccessNum *int64 `json:"success_num,omitempty"`

	// 处理时间
	ProcessTime *int64 `json:"process_time,omitempty"`

	// 调用时间
	InvokeTime *int64 `json:"invoke_time,omitempty"`
}

func (o EventSubMetricsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventSubMetricsItem struct{}"
	}

	return strings.Join([]string{"EventSubMetricsItem", string(data)}, " ")
}
