package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QueuePlan 查询指定队列定时扩缩容信息的响应参数。
type QueuePlan struct {

	// 扩缩容计划的ID编号。
	Id *int64 `json:"id,omitempty"`

	// 扩缩容计划的名称。
	PlanName *string `json:"plan_name,omitempty"`

	// 队列扩缩容计划的目标CU值。
	TargetCu *int32 `json:"target_cu,omitempty"`

	// 队列扩缩容计划的起始时（24小时制）。
	StartHour *int32 `json:"start_hour,omitempty"`

	// 定时扩缩容计划的起始分。
	StartMinute *int32 `json:"start_minute,omitempty"`

	// 定时扩缩容计划的重复周期规律
	RepeatDay []QueuePlanRepeatDay `json:"repeat_day"`

	// 有效期开始时间（13位时间戳）
	ValidDateBegin *int64 `json:"valid_date_begin,omitempty"`

	// 有效期结束时间（13位时间戳）
	ValidDateEnd *int64 `json:"valid_date_end,omitempty"`

	// 当前的扩缩容计划是否激活。
	Activate *bool `json:"activate,omitempty"`

	// 当前扩缩容计划最近一次执行的时间。
	LastExecuteTime *int64 `json:"last_execute_time,omitempty"`
}

func (o QueuePlan) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueuePlan struct{}"
	}

	return strings.Join([]string{"QueuePlan", string(data)}, " ")
}

type QueuePlanRepeatDay struct {
	value string
}

type QueuePlanRepeatDayEnum struct {
	SUNDAY    QueuePlanRepeatDay
	MONDAY    QueuePlanRepeatDay
	TUESDAY   QueuePlanRepeatDay
	WEDNESDAY QueuePlanRepeatDay
	THURSDAY  QueuePlanRepeatDay
	FRIDAY    QueuePlanRepeatDay
	SATURDAY  QueuePlanRepeatDay
}

func GetQueuePlanRepeatDayEnum() QueuePlanRepeatDayEnum {
	return QueuePlanRepeatDayEnum{
		SUNDAY: QueuePlanRepeatDay{
			value: "SUNDAY",
		},
		MONDAY: QueuePlanRepeatDay{
			value: "MONDAY",
		},
		TUESDAY: QueuePlanRepeatDay{
			value: "TUESDAY",
		},
		WEDNESDAY: QueuePlanRepeatDay{
			value: "WEDNESDAY",
		},
		THURSDAY: QueuePlanRepeatDay{
			value: "THURSDAY",
		},
		FRIDAY: QueuePlanRepeatDay{
			value: "FRIDAY",
		},
		SATURDAY: QueuePlanRepeatDay{
			value: "SATURDAY",
		},
	}
}

func (c QueuePlanRepeatDay) Value() string {
	return c.value
}

func (c QueuePlanRepeatDay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueuePlanRepeatDay) UnmarshalJSON(b []byte) error {
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
