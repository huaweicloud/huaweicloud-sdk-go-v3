package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConditionInfo 剧本触发规则详情
type ConditionInfo struct {

	// 表达式类型。默认为common,事件触发剧本必填
	ExpressionType *string `json:"expression_type,omitempty"`

	// 触发条件。事件触发剧本必填
	Conditions *[]ConditionItem `json:"conditions,omitempty"`

	// 条件逻辑组合。事件触发剧本必填
	Logics *[]string `json:"logics,omitempty"`

	// Cron 表达式（定时任务）。定时触发剧本必填
	Cron *string `json:"cron,omitempty"`

	// 定时重复类型(second--秒, hour--小时,day--天，week-周)。定时触发剧本必填
	ScheduleType *string `json:"schedule_type,omitempty"`

	// 剧本开始执行类型，IMMEDIATELY--创建完成立即执行, CUSTOM--自定义执行时间。定时触发剧本必填
	StartType *string `json:"start_type,omitempty"`

	// 剧本结束执行类型，FOREVER--一直执行, CUSTOM--自定义结束时间。定时触发剧本必填
	EndType *string `json:"end_type,omitempty"`

	// 定时结束时间。定时触发剧本必填
	EndTime *string `json:"end_time,omitempty"`

	// 执行时间段 2021-01-30T23:00:00Z+0800。定时触发剧本必填
	RepeatRange *string `json:"repeat_range,omitempty"`

	// 是否只执行一次。定时触发剧本必填
	OnlyOnce *bool `json:"only_once,omitempty"`

	// 执行队列类型（PARALLEL-新任务与之前任务并行）。定时触发剧本必填
	ExecutionType *string `json:"execution_type,omitempty"`
}

func (o ConditionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditionInfo struct{}"
	}

	return strings.Join([]string{"ConditionInfo", string(data)}, " ")
}
