package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type FrequencyRespBody struct {

	// 时间类型。
	Type *FrequencyRespBodyType `json:"type,omitempty"`

	// 当字段type为\"CRON\"时取该字段。
	CronExpr *string `json:"cron_expr,omitempty"`

	// 当字段type为\"DAILY\"或者\"WEEKLY\"时取该字段。
	HourOfDay *int32 `json:"hour_of_day,omitempty"`

	// 当字段type为\"WEEKLY\"时取该字段（周日~周六）。
	DayOfWeek *int32 `json:"day_of_week,omitempty"`

	// 当字段type为\"FIXED_RATE\"时取该字段（当fixed_rate_unit单位为minute，最大值60；当fixed_rate_unit单位为hour，最大值24）。
	FixedRate *int32 `json:"fixed_rate,omitempty"`

	// 时间单位枚举值：
	FixedRateUnit *FrequencyRespBodyFixedRateUnit `json:"fixed_rate_unit,omitempty"`
}

func (o FrequencyRespBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FrequencyRespBody struct{}"
	}

	return strings.Join([]string{"FrequencyRespBody", string(data)}, " ")
}

type FrequencyRespBodyType struct {
	value string
}

type FrequencyRespBodyTypeEnum struct {
	CRON       FrequencyRespBodyType
	HOURLY     FrequencyRespBodyType
	DAILY      FrequencyRespBodyType
	WEEKLY     FrequencyRespBodyType
	FIXED_RATE FrequencyRespBodyType
}

func GetFrequencyRespBodyTypeEnum() FrequencyRespBodyTypeEnum {
	return FrequencyRespBodyTypeEnum{
		CRON: FrequencyRespBodyType{
			value: "CRON",
		},
		HOURLY: FrequencyRespBodyType{
			value: "HOURLY",
		},
		DAILY: FrequencyRespBodyType{
			value: "DAILY",
		},
		WEEKLY: FrequencyRespBodyType{
			value: "WEEKLY",
		},
		FIXED_RATE: FrequencyRespBodyType{
			value: "FIXED_RATE",
		},
	}
}

func (c FrequencyRespBodyType) Value() string {
	return c.value
}

func (c FrequencyRespBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FrequencyRespBodyType) UnmarshalJSON(b []byte) error {
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

type FrequencyRespBodyFixedRateUnit struct {
	value string
}

type FrequencyRespBodyFixedRateUnitEnum struct {
	MINUTE FrequencyRespBodyFixedRateUnit
	HOUR   FrequencyRespBodyFixedRateUnit
}

func GetFrequencyRespBodyFixedRateUnitEnum() FrequencyRespBodyFixedRateUnitEnum {
	return FrequencyRespBodyFixedRateUnitEnum{
		MINUTE: FrequencyRespBodyFixedRateUnit{
			value: "minute",
		},
		HOUR: FrequencyRespBodyFixedRateUnit{
			value: "hour",
		},
	}
}

func (c FrequencyRespBodyFixedRateUnit) Value() string {
	return c.value
}

func (c FrequencyRespBodyFixedRateUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FrequencyRespBodyFixedRateUnit) UnmarshalJSON(b []byte) error {
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
