package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QueuePlanEntity 查询指定队列定时扩缩容信息的响应参数。
type QueuePlanEntity struct {

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
	RepeatDay []QueuePlanEntityRepeatDay `json:"repeat_day"`

	// 有效期开始时间（13位时间戳）
	ValidDateBegin *int64 `json:"valid_date_begin,omitempty"`

	// 有效期结束时间（13位时间戳）
	ValidDateEnd *int64 `json:"valid_date_end,omitempty"`

	// 当前的扩缩容计划是否激活。
	Activate *bool `json:"activate,omitempty"`

	// 当前扩缩容计划最近一次执行的时间。
	LastExecuteTime *int64 `json:"last_execute_time,omitempty"`
}

func (o QueuePlanEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueuePlanEntity struct{}"
	}

	return strings.Join([]string{"QueuePlanEntity", string(data)}, " ")
}

type QueuePlanEntityRepeatDay struct {
	value string
}

type QueuePlanEntityRepeatDayEnum struct {
	SUNDAY    QueuePlanEntityRepeatDay
	MONDAY    QueuePlanEntityRepeatDay
	TUESDAY   QueuePlanEntityRepeatDay
	WEDNESDAY QueuePlanEntityRepeatDay
	THURSDAY  QueuePlanEntityRepeatDay
	FRIDAY    QueuePlanEntityRepeatDay
	SATURDAY  QueuePlanEntityRepeatDay
}

func GetQueuePlanEntityRepeatDayEnum() QueuePlanEntityRepeatDayEnum {
	return QueuePlanEntityRepeatDayEnum{
		SUNDAY: QueuePlanEntityRepeatDay{
			value: "SUNDAY",
		},
		MONDAY: QueuePlanEntityRepeatDay{
			value: "MONDAY",
		},
		TUESDAY: QueuePlanEntityRepeatDay{
			value: "TUESDAY",
		},
		WEDNESDAY: QueuePlanEntityRepeatDay{
			value: "WEDNESDAY",
		},
		THURSDAY: QueuePlanEntityRepeatDay{
			value: "THURSDAY",
		},
		FRIDAY: QueuePlanEntityRepeatDay{
			value: "FRIDAY",
		},
		SATURDAY: QueuePlanEntityRepeatDay{
			value: "SATURDAY",
		},
	}
}

func (c QueuePlanEntityRepeatDay) Value() string {
	return c.value
}

func (c QueuePlanEntityRepeatDay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueuePlanEntityRepeatDay) UnmarshalJSON(b []byte) error {
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
