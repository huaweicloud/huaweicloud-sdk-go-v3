package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QueuePlanRequestBody 创建扩缩容计划的body体。
type QueuePlanRequestBody struct {

	// 队列扩缩容计划名称，名称只能包含数字、英文字母和下划线，但不能是纯数字，且不能以下划线开头。
	PlanName string `json:"plan_name"`

	// 队列扩缩容计划CU的目标值
	TargetCu int32 `json:"target_cu"`

	// 队列扩缩容计划起始小时时间
	StartHour int32 `json:"start_hour"`

	// 队列扩缩容计划的起始分钟时间
	StartMinute int32 `json:"start_minute"`

	// 定时扩缩容计划的重复周期规律，可以选择周一到周日的某一天、某几天、或者不选择。如果不选择，则会在当前时间后的start_hour：start_minute时间点执行扩缩容计划。如：\"repeat_day\": [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\",\"SUNDAY\"]
	RepeatDay []QueuePlanRequestBodyRepeatDay `json:"repeat_day"`

	// 有效期开始时间（13位时间戳）
	ValidDateBegin *int64 `json:"valid_date_begin,omitempty"`

	// 有效期结束时间（13位时间戳）
	ValidDateEnd *int64 `json:"valid_date_end,omitempty"`

	// 当前设置的扩缩容计划是否激活，默认为激活
	Activate *bool `json:"activate,omitempty"`
}

func (o QueuePlanRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueuePlanRequestBody struct{}"
	}

	return strings.Join([]string{"QueuePlanRequestBody", string(data)}, " ")
}

type QueuePlanRequestBodyRepeatDay struct {
	value string
}

type QueuePlanRequestBodyRepeatDayEnum struct {
	SUNDAY    QueuePlanRequestBodyRepeatDay
	MONDAY    QueuePlanRequestBodyRepeatDay
	TUESDAY   QueuePlanRequestBodyRepeatDay
	WEDNESDAY QueuePlanRequestBodyRepeatDay
	THURSDAY  QueuePlanRequestBodyRepeatDay
	FRIDAY    QueuePlanRequestBodyRepeatDay
	SATURDAY  QueuePlanRequestBodyRepeatDay
}

func GetQueuePlanRequestBodyRepeatDayEnum() QueuePlanRequestBodyRepeatDayEnum {
	return QueuePlanRequestBodyRepeatDayEnum{
		SUNDAY: QueuePlanRequestBodyRepeatDay{
			value: "SUNDAY",
		},
		MONDAY: QueuePlanRequestBodyRepeatDay{
			value: "MONDAY",
		},
		TUESDAY: QueuePlanRequestBodyRepeatDay{
			value: "TUESDAY",
		},
		WEDNESDAY: QueuePlanRequestBodyRepeatDay{
			value: "WEDNESDAY",
		},
		THURSDAY: QueuePlanRequestBodyRepeatDay{
			value: "THURSDAY",
		},
		FRIDAY: QueuePlanRequestBodyRepeatDay{
			value: "FRIDAY",
		},
		SATURDAY: QueuePlanRequestBodyRepeatDay{
			value: "SATURDAY",
		},
	}
}

func (c QueuePlanRequestBodyRepeatDay) Value() string {
	return c.value
}

func (c QueuePlanRequestBodyRepeatDay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueuePlanRequestBodyRepeatDay) UnmarshalJSON(b []byte) error {
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
