package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Schedule struct {

	// frequency_interval
	FrequencyInterval int32 `json:"frequency_interval"`

	// frequency_unit. MINUTE, HOUR, DAY.
	FrequencyUnit ScheduleFrequencyUnit `json:"frequency_unit"`

	// frequency_interval
	PeriodInterval int32 `json:"period_interval"`

	// period_unit. MINUTE, HOUR, DAY.
	PeriodUnit SchedulePeriodUnit `json:"period_unit"`

	// delay_interval
	DelayInterval *int32 `json:"delay_interval,omitempty"`

	// overtime_interval
	OvertimeInterval *int32 `json:"overtime_interval,omitempty"`
}

func (o Schedule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Schedule struct{}"
	}

	return strings.Join([]string{"Schedule", string(data)}, " ")
}

type ScheduleFrequencyUnit struct {
	value string
}

type ScheduleFrequencyUnitEnum struct {
	MINUTE ScheduleFrequencyUnit
	HOUR   ScheduleFrequencyUnit
	DAY    ScheduleFrequencyUnit
}

func GetScheduleFrequencyUnitEnum() ScheduleFrequencyUnitEnum {
	return ScheduleFrequencyUnitEnum{
		MINUTE: ScheduleFrequencyUnit{
			value: "MINUTE",
		},
		HOUR: ScheduleFrequencyUnit{
			value: "HOUR",
		},
		DAY: ScheduleFrequencyUnit{
			value: "DAY",
		},
	}
}

func (c ScheduleFrequencyUnit) Value() string {
	return c.value
}

func (c ScheduleFrequencyUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduleFrequencyUnit) UnmarshalJSON(b []byte) error {
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

type SchedulePeriodUnit struct {
	value string
}

type SchedulePeriodUnitEnum struct {
	MINUTE SchedulePeriodUnit
	HOUR   SchedulePeriodUnit
	DAY    SchedulePeriodUnit
}

func GetSchedulePeriodUnitEnum() SchedulePeriodUnitEnum {
	return SchedulePeriodUnitEnum{
		MINUTE: SchedulePeriodUnit{
			value: "MINUTE",
		},
		HOUR: SchedulePeriodUnit{
			value: "HOUR",
		},
		DAY: SchedulePeriodUnit{
			value: "DAY",
		},
	}
}

func (c SchedulePeriodUnit) Value() string {
	return c.value
}

func (c SchedulePeriodUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SchedulePeriodUnit) UnmarshalJSON(b []byte) error {
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
