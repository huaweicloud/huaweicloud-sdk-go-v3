package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateSqlAlarmRuleFrequency struct {

	// 时间类型。
	Type CreateSqlAlarmRuleFrequencyType `json:"type"`

	// 当字段type为\"CRON\"时取该字段。
	CronExpr *string `json:"cron_expr,omitempty"`

	// 当字段type为\"DAILY\"或者\"WEEKLY\"时取该字段。  DAILY：最小值：0，最大值：23 WEEKLY：最小值：0，最大值：23
	HourOfDay *int32 `json:"hour_of_day,omitempty"`

	// 当字段type为\"WEEKLY\"时取该字段（周日~周六）。
	DayOfWeek *int32 `json:"day_of_week,omitempty"`

	// 当字段type为\"FIXED_RATE\"时取该字段（当fixed_rate_unit单位为minute，最大值60；当fixed_rate_unit单位为hour，最大值24）。
	FixedRate *int32 `json:"fixed_rate,omitempty"`

	// 时间单位。
	FixedRateUnit *CreateSqlAlarmRuleFrequencyFixedRateUnit `json:"fixed_rate_unit,omitempty"`
}

func (o CreateSqlAlarmRuleFrequency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlAlarmRuleFrequency struct{}"
	}

	return strings.Join([]string{"CreateSqlAlarmRuleFrequency", string(data)}, " ")
}

type CreateSqlAlarmRuleFrequencyType struct {
	value string
}

type CreateSqlAlarmRuleFrequencyTypeEnum struct {
	CRON       CreateSqlAlarmRuleFrequencyType
	HOURLY     CreateSqlAlarmRuleFrequencyType
	DAILY      CreateSqlAlarmRuleFrequencyType
	WEEKLY     CreateSqlAlarmRuleFrequencyType
	FIXED_RATE CreateSqlAlarmRuleFrequencyType
}

func GetCreateSqlAlarmRuleFrequencyTypeEnum() CreateSqlAlarmRuleFrequencyTypeEnum {
	return CreateSqlAlarmRuleFrequencyTypeEnum{
		CRON: CreateSqlAlarmRuleFrequencyType{
			value: "CRON",
		},
		HOURLY: CreateSqlAlarmRuleFrequencyType{
			value: "HOURLY",
		},
		DAILY: CreateSqlAlarmRuleFrequencyType{
			value: "DAILY",
		},
		WEEKLY: CreateSqlAlarmRuleFrequencyType{
			value: "WEEKLY",
		},
		FIXED_RATE: CreateSqlAlarmRuleFrequencyType{
			value: "FIXED_RATE",
		},
	}
}

func (c CreateSqlAlarmRuleFrequencyType) Value() string {
	return c.value
}

func (c CreateSqlAlarmRuleFrequencyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSqlAlarmRuleFrequencyType) UnmarshalJSON(b []byte) error {
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

type CreateSqlAlarmRuleFrequencyFixedRateUnit struct {
	value string
}

type CreateSqlAlarmRuleFrequencyFixedRateUnitEnum struct {
	MINUTE CreateSqlAlarmRuleFrequencyFixedRateUnit
	HOUR   CreateSqlAlarmRuleFrequencyFixedRateUnit
}

func GetCreateSqlAlarmRuleFrequencyFixedRateUnitEnum() CreateSqlAlarmRuleFrequencyFixedRateUnitEnum {
	return CreateSqlAlarmRuleFrequencyFixedRateUnitEnum{
		MINUTE: CreateSqlAlarmRuleFrequencyFixedRateUnit{
			value: "minute",
		},
		HOUR: CreateSqlAlarmRuleFrequencyFixedRateUnit{
			value: "hour",
		},
	}
}

func (c CreateSqlAlarmRuleFrequencyFixedRateUnit) Value() string {
	return c.value
}

func (c CreateSqlAlarmRuleFrequencyFixedRateUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSqlAlarmRuleFrequencyFixedRateUnit) UnmarshalJSON(b []byte) error {
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
