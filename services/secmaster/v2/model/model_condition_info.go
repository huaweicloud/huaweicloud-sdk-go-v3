package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Condition info of rule
type ConditionInfo struct {

	// expression type, all, any, user_define
	ExpressionType *string `json:"expression_type,omitempty"`

	// Information of conditions.
	Conditions *[]ConditionItem `json:"conditions,omitempty"`

	// Logic item of condition.
	Logics *interface{} `json:"logics,omitempty"`

	// Cron 表达式
	Cron *string `json:"cron,omitempty"`

	// schedule type, second hours...
	ScheduleType *string `json:"schedule_type,omitempty"`

	// 执行时间段 2021-01-30T23:00:00Z+0800
	RepeatRange *string `json:"repeat_range,omitempty"`

	// 重复次数
	RepeatCount *string `json:"repeat_count,omitempty"`
}

func (o ConditionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditionInfo struct{}"
	}

	return strings.Join([]string{"ConditionInfo", string(data)}, " ")
}
