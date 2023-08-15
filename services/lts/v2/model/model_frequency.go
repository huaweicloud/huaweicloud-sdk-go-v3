package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Frequency struct {

	// 时间类型
	Type FrequencyType `json:"type"`

	// 当字段type为\"CRON\"时取该字段
	CronExpr *string `json:"cron_expr,omitempty"`

	// 当字段type为\"DAILY\"或者\"WEEKLY\"时取该字段
	HourOfDay *int32 `json:"hour_of_day,omitempty"`

	// 当字段type为\"WEEKLY\"时取该字段(周日~周六)
	DayOfWeek *int32 `json:"day_of_week,omitempty"`

	// 当字段type为\"FIXED_RATE\"时取该字段(当fixed_rate_unit单位为minute，最大值60;当fixed_rate_unit单位为hour，最大值24)
	FixedRate *int32 `json:"fixed_rate,omitempty"`

	// 时间单位
	FixedRateUnit *FrequencyFixedRateUnit `json:"fixed_rate_unit,omitempty"`
}

func (o Frequency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Frequency struct{}"
	}

	return strings.Join([]string{"Frequency", string(data)}, " ")
}

type FrequencyType struct {
	value string
}

type FrequencyTypeEnum struct {
	CRON       FrequencyType
	HOURLY     FrequencyType
	DAILY      FrequencyType
	WEEKLY     FrequencyType
	FIXED_RATE FrequencyType
}

func GetFrequencyTypeEnum() FrequencyTypeEnum {
	return FrequencyTypeEnum{
		CRON: FrequencyType{
			value: "CRON",
		},
		HOURLY: FrequencyType{
			value: "HOURLY",
		},
		DAILY: FrequencyType{
			value: "DAILY",
		},
		WEEKLY: FrequencyType{
			value: "WEEKLY",
		},
		FIXED_RATE: FrequencyType{
			value: "FIXED_RATE",
		},
	}
}

func (c FrequencyType) Value() string {
	return c.value
}

func (c FrequencyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FrequencyType) UnmarshalJSON(b []byte) error {
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

type FrequencyFixedRateUnit struct {
	value string
}

type FrequencyFixedRateUnitEnum struct {
	MINUTE FrequencyFixedRateUnit
	HOUR   FrequencyFixedRateUnit
}

func GetFrequencyFixedRateUnitEnum() FrequencyFixedRateUnitEnum {
	return FrequencyFixedRateUnitEnum{
		MINUTE: FrequencyFixedRateUnit{
			value: "minute",
		},
		HOUR: FrequencyFixedRateUnit{
			value: "hour",
		},
	}
}

func (c FrequencyFixedRateUnit) Value() string {
	return c.value
}

func (c FrequencyFixedRateUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FrequencyFixedRateUnit) UnmarshalJSON(b []byte) error {
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
