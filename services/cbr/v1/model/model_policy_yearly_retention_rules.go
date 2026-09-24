package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyYearlyRetentionRules 设置年备备份的保留规则
type PolicyYearlyRetentionRules struct {

	// 年备规则的类型
	RetentionType *PolicyYearlyRetentionRulesRetentionType `json:"retention_type,omitempty"`

	// 将每年中指定月份的备份设置为年备备份，当retention_type为Weekly时，需要与retention_weeks和days_of_week共同设置，当retention_type为Monthly时，需要与days_of_month共同设置
	RetentionMonths *[]PolicyYearlyRetentionRulesRetentionMonths `json:"retention_months,omitempty"`

	// 将选中月份的第几个星期的备份设置为年备备份，当retention_type为Weekly时才能设置，设置时需要与retention_months和days_of_week共同设置
	RetentionWeeks *[]PolicyYearlyRetentionRulesRetentionWeeks `json:"retention_weeks,omitempty"`

	// 表示将选中月份的指定天的备份设置为年备备份，当retention_type为Monthly时才能设置，取值范围为1-28和-1，-1代表每个月的最后一天，需要与retention_months共同设置
	DaysOfMonth *[]int32 `json:"days_of_month,omitempty"`

	// 设置指定月份的指定星期中的指定天的备份为年备备份，当retention_type为Weekly时才能设置，设置时需要与retention_weeks和retention_months共同设置
	DaysOfWeek *[]PolicyYearlyRetentionRulesDaysOfWeek `json:"days_of_week,omitempty"`

	// 年备备份的保留时间，取值范围为1-100，以及-1，单位为年，-1代表年备策略不启用
	RetentionDurationPeriods *int32 `json:"retention_duration_periods,omitempty"`
}

func (o PolicyYearlyRetentionRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyYearlyRetentionRules struct{}"
	}

	return strings.Join([]string{"PolicyYearlyRetentionRules", string(data)}, " ")
}

type PolicyYearlyRetentionRulesRetentionType struct {
	value string
}

type PolicyYearlyRetentionRulesRetentionTypeEnum struct {
	WEEKLYMONTHLY PolicyYearlyRetentionRulesRetentionType
}

func GetPolicyYearlyRetentionRulesRetentionTypeEnum() PolicyYearlyRetentionRulesRetentionTypeEnum {
	return PolicyYearlyRetentionRulesRetentionTypeEnum{
		WEEKLYMONTHLY: PolicyYearlyRetentionRulesRetentionType{
			value: "WEEKLY，MONTHLY",
		},
	}
}

func (c PolicyYearlyRetentionRulesRetentionType) Value() string {
	return c.value
}

func (c PolicyYearlyRetentionRulesRetentionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyYearlyRetentionRulesRetentionType) UnmarshalJSON(b []byte) error {
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

type PolicyYearlyRetentionRulesRetentionMonths struct {
	value string
}

type PolicyYearlyRetentionRulesRetentionMonthsEnum struct {
	JANUARY   PolicyYearlyRetentionRulesRetentionMonths
	FEBRUARY  PolicyYearlyRetentionRulesRetentionMonths
	MARCH     PolicyYearlyRetentionRulesRetentionMonths
	APRIL     PolicyYearlyRetentionRulesRetentionMonths
	MAY       PolicyYearlyRetentionRulesRetentionMonths
	JUNE      PolicyYearlyRetentionRulesRetentionMonths
	JULY      PolicyYearlyRetentionRulesRetentionMonths
	AUGUST    PolicyYearlyRetentionRulesRetentionMonths
	SEPTEMBER PolicyYearlyRetentionRulesRetentionMonths
	OCTOBER   PolicyYearlyRetentionRulesRetentionMonths
	NOVEMBER  PolicyYearlyRetentionRulesRetentionMonths
	DECEMBER  PolicyYearlyRetentionRulesRetentionMonths
}

func GetPolicyYearlyRetentionRulesRetentionMonthsEnum() PolicyYearlyRetentionRulesRetentionMonthsEnum {
	return PolicyYearlyRetentionRulesRetentionMonthsEnum{
		JANUARY: PolicyYearlyRetentionRulesRetentionMonths{
			value: "JANUARY",
		},
		FEBRUARY: PolicyYearlyRetentionRulesRetentionMonths{
			value: "FEBRUARY",
		},
		MARCH: PolicyYearlyRetentionRulesRetentionMonths{
			value: "MARCH",
		},
		APRIL: PolicyYearlyRetentionRulesRetentionMonths{
			value: "APRIL",
		},
		MAY: PolicyYearlyRetentionRulesRetentionMonths{
			value: "MAY",
		},
		JUNE: PolicyYearlyRetentionRulesRetentionMonths{
			value: "JUNE",
		},
		JULY: PolicyYearlyRetentionRulesRetentionMonths{
			value: "JULY",
		},
		AUGUST: PolicyYearlyRetentionRulesRetentionMonths{
			value: "AUGUST",
		},
		SEPTEMBER: PolicyYearlyRetentionRulesRetentionMonths{
			value: "SEPTEMBER",
		},
		OCTOBER: PolicyYearlyRetentionRulesRetentionMonths{
			value: "OCTOBER",
		},
		NOVEMBER: PolicyYearlyRetentionRulesRetentionMonths{
			value: "NOVEMBER",
		},
		DECEMBER: PolicyYearlyRetentionRulesRetentionMonths{
			value: "DECEMBER",
		},
	}
}

func (c PolicyYearlyRetentionRulesRetentionMonths) Value() string {
	return c.value
}

func (c PolicyYearlyRetentionRulesRetentionMonths) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyYearlyRetentionRulesRetentionMonths) UnmarshalJSON(b []byte) error {
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

type PolicyYearlyRetentionRulesRetentionWeeks struct {
	value string
}

type PolicyYearlyRetentionRulesRetentionWeeksEnum struct {
	FIRST  PolicyYearlyRetentionRulesRetentionWeeks
	SECOND PolicyYearlyRetentionRulesRetentionWeeks
	THIRD  PolicyYearlyRetentionRulesRetentionWeeks
	FOURTH PolicyYearlyRetentionRulesRetentionWeeks
	LAST   PolicyYearlyRetentionRulesRetentionWeeks
}

func GetPolicyYearlyRetentionRulesRetentionWeeksEnum() PolicyYearlyRetentionRulesRetentionWeeksEnum {
	return PolicyYearlyRetentionRulesRetentionWeeksEnum{
		FIRST: PolicyYearlyRetentionRulesRetentionWeeks{
			value: "FIRST",
		},
		SECOND: PolicyYearlyRetentionRulesRetentionWeeks{
			value: "SECOND",
		},
		THIRD: PolicyYearlyRetentionRulesRetentionWeeks{
			value: "THIRD",
		},
		FOURTH: PolicyYearlyRetentionRulesRetentionWeeks{
			value: "FOURTH",
		},
		LAST: PolicyYearlyRetentionRulesRetentionWeeks{
			value: "LAST",
		},
	}
}

func (c PolicyYearlyRetentionRulesRetentionWeeks) Value() string {
	return c.value
}

func (c PolicyYearlyRetentionRulesRetentionWeeks) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyYearlyRetentionRulesRetentionWeeks) UnmarshalJSON(b []byte) error {
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

type PolicyYearlyRetentionRulesDaysOfWeek struct {
	value string
}

type PolicyYearlyRetentionRulesDaysOfWeekEnum struct {
	MO PolicyYearlyRetentionRulesDaysOfWeek
	TU PolicyYearlyRetentionRulesDaysOfWeek
	WE PolicyYearlyRetentionRulesDaysOfWeek
	TH PolicyYearlyRetentionRulesDaysOfWeek
	FR PolicyYearlyRetentionRulesDaysOfWeek
	SA PolicyYearlyRetentionRulesDaysOfWeek
	SU PolicyYearlyRetentionRulesDaysOfWeek
}

func GetPolicyYearlyRetentionRulesDaysOfWeekEnum() PolicyYearlyRetentionRulesDaysOfWeekEnum {
	return PolicyYearlyRetentionRulesDaysOfWeekEnum{
		MO: PolicyYearlyRetentionRulesDaysOfWeek{
			value: "MO",
		},
		TU: PolicyYearlyRetentionRulesDaysOfWeek{
			value: "TU",
		},
		WE: PolicyYearlyRetentionRulesDaysOfWeek{
			value: "WE",
		},
		TH: PolicyYearlyRetentionRulesDaysOfWeek{
			value: "TH",
		},
		FR: PolicyYearlyRetentionRulesDaysOfWeek{
			value: "FR",
		},
		SA: PolicyYearlyRetentionRulesDaysOfWeek{
			value: "SA",
		},
		SU: PolicyYearlyRetentionRulesDaysOfWeek{
			value: "SU",
		},
	}
}

func (c PolicyYearlyRetentionRulesDaysOfWeek) Value() string {
	return c.value
}

func (c PolicyYearlyRetentionRulesDaysOfWeek) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyYearlyRetentionRulesDaysOfWeek) UnmarshalJSON(b []byte) error {
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
