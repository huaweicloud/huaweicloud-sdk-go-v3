package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TriggerTime 定时任务触发时间
type TriggerTime struct {

	// 时区
	TimeZone string `json:"time_zone"`

	// 策略
	Policy TriggerTimePolicy `json:"policy"`

	// 单次执行的执行时间戳
	SingleScheduledTime *int64 `json:"single_scheduled_time,omitempty"`

	// 周期执行的执行当天的时间 \"00:00:00\"
	PeriodicScheduledTime *string `json:"periodic_scheduled_time,omitempty"`

	// 周期执行时的具体星期列表按逗号分割, 如星期日为\"1\",星期一为\"2\"
	Period *string `json:"period,omitempty"`

	// cron表达式
	Cron *string `json:"cron,omitempty"`

	// 定时任务规则截止日期时间戳
	ScheduledCloseTime *int64 `json:"scheduled_close_time,omitempty"`
}

func (o TriggerTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TriggerTime struct{}"
	}

	return strings.Join([]string{"TriggerTime", string(data)}, " ")
}

type TriggerTimePolicy struct {
	value string
}

type TriggerTimePolicyEnum struct {
	PERIODIC TriggerTimePolicy
	ONCE     TriggerTimePolicy
	CRON     TriggerTimePolicy
}

func GetTriggerTimePolicyEnum() TriggerTimePolicyEnum {
	return TriggerTimePolicyEnum{
		PERIODIC: TriggerTimePolicy{
			value: "PERIODIC",
		},
		ONCE: TriggerTimePolicy{
			value: "ONCE",
		},
		CRON: TriggerTimePolicy{
			value: "CRON",
		},
	}
}

func (c TriggerTimePolicy) Value() string {
	return c.value
}

func (c TriggerTimePolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerTimePolicy) UnmarshalJSON(b []byte) error {
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
