package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 计划任务的配置项，选填。计划任务不支持obs输入，对于url输入则仅支持rtmp和rtsp两种形式。
type TaskTiming struct {
	// 计划任务的类型，使用计划任务时必填。可选类型分别为once（仅执行一次），daily（每日执行），weekly（每周执行），monthly（每月执行）。

	Type TaskTimingType `json:"type"`
	// 用户所处的时区，使用计划任务时必填。精确到分钟。

	Timezone string `json:"timezone"`
	// 作业会在一周的哪几天执行，当且仅当计划任务类型为weekly时，该字段需填且必填。1~7分别指代星期一至星期日。

	DaysOfWeek *[]int32 `json:"days_of_week,omitempty"`
	// 作业会在一个月的哪几天执行，当且仅当计划任务类型为monthly时，该字段需填且必填。1~31分别指代一个月中的1日至31日。

	DaysOfMonth *[]int32 `json:"days_of_month,omitempty"`
	// 作业的执行日。当且仅当计划任务类型为once且为频率模式时，该字段需填且必填。格式形如yyyy-MM-dd。

	Date *string `json:"date,omitempty"`
	// 时间段模式配置。和frequency字段二选一，不可共存。时间段模式下，至少需指定一个时间段。

	Periods *[]TaskTimingPeriods `json:"periods,omitempty"`

	Frequency *TaskTimingFrequency `json:"frequency,omitempty"`
}

func (o TaskTiming) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskTiming struct{}"
	}

	return strings.Join([]string{"TaskTiming", string(data)}, " ")
}

type TaskTimingType struct {
	value string
}

type TaskTimingTypeEnum struct {
	ONCE    TaskTimingType
	DAILY   TaskTimingType
	WEEKLY  TaskTimingType
	MONTHLY TaskTimingType
}

func GetTaskTimingTypeEnum() TaskTimingTypeEnum {
	return TaskTimingTypeEnum{
		ONCE: TaskTimingType{
			value: "once",
		},
		DAILY: TaskTimingType{
			value: "daily",
		},
		WEEKLY: TaskTimingType{
			value: "weekly",
		},
		MONTHLY: TaskTimingType{
			value: "monthly",
		},
	}
}

func (c TaskTimingType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskTimingType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
