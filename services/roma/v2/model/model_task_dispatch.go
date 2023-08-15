package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TaskDispatch 调度计划详情
type TaskDispatch struct {

	// 调度计划的执行开始时间
	StartDatetime *int64 `json:"start_datetime,omitempty"`

	// 调度计划执行周期的时间单位，当使用cron表达式时，为空 - MIN (分钟) - HOUR (小时) - DAY (日) - WEEK (周) - MON (月)
	Period *TaskDispatchPeriod `json:"period,omitempty"`

	// 调度计划的执行间隔时间周期
	DispatchInterval *int64 `json:"dispatch_interval,omitempty"`

	// 调度计划的备注信息
	Remark *string `json:"remark,omitempty"`

	// 调度计划是否使用cron表达式，允许如下值： - true (使用cron表达式) - false (不使用cron表达式)
	UseQuartzCron *bool `json:"use_quartz_cron,omitempty"`

	// 调度计划的cron表达式
	Cron *string `json:"cron,omitempty"`
}

func (o TaskDispatch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskDispatch struct{}"
	}

	return strings.Join([]string{"TaskDispatch", string(data)}, " ")
}

type TaskDispatchPeriod struct {
	value string
}

type TaskDispatchPeriodEnum struct {
	MIN  TaskDispatchPeriod
	HOUR TaskDispatchPeriod
	DAY  TaskDispatchPeriod
	WEEK TaskDispatchPeriod
	MON  TaskDispatchPeriod
}

func GetTaskDispatchPeriodEnum() TaskDispatchPeriodEnum {
	return TaskDispatchPeriodEnum{
		MIN: TaskDispatchPeriod{
			value: "MIN",
		},
		HOUR: TaskDispatchPeriod{
			value: "HOUR",
		},
		DAY: TaskDispatchPeriod{
			value: "DAY",
		},
		WEEK: TaskDispatchPeriod{
			value: "WEEK",
		},
		MON: TaskDispatchPeriod{
			value: "MON",
		},
	}
}

func (c TaskDispatchPeriod) Value() string {
	return c.value
}

func (c TaskDispatchPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskDispatchPeriod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
