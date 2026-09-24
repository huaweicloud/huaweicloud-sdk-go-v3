package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PolicyMonthlyRetentionRules struct {

	// 月备规则的类型
	RetentionType *PolicyMonthlyRetentionRulesRetentionType `json:"retention_type,omitempty"`

	// 将每月第几个星期的备份设置为月备备份，当retention_type为Weekly时才能设置，设置时需要与days_of_week共同设置
	RetentionWeeks *[]PolicyMonthlyRetentionRulesRetentionWeeks `json:"retention_weeks,omitempty"`

	// 设置选中的星期中的指定天的备份为月备备份，当retention_type为Weekly时才能设置，设置时需要与retention_weeks共同设置
	DaysOfWeek *[]PolicyMonthlyRetentionRulesDaysOfWeek `json:"days_of_week,omitempty"`

	// 表示将每个月中的指定天设置为月备备份，当retention_type为Monthly时才能设置，取值范围为1-28和-1，-1代表每个月的最后一天
	DaysOfMonth *[]int32 `json:"days_of_month,omitempty"`

	// 月备备份的保留时间，取值范围为1-1200，以及-1，单位为月，-1代表月备策略不启用
	RetentionDurationPeriods *int32 `json:"retention_duration_periods,omitempty"`
}

func (o PolicyMonthlyRetentionRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyMonthlyRetentionRules struct{}"
	}

	return strings.Join([]string{"PolicyMonthlyRetentionRules", string(data)}, " ")
}

type PolicyMonthlyRetentionRulesRetentionType struct {
	value string
}

type PolicyMonthlyRetentionRulesRetentionTypeEnum struct {
	WEEKLY  PolicyMonthlyRetentionRulesRetentionType
	MONTHLY PolicyMonthlyRetentionRulesRetentionType
}

func GetPolicyMonthlyRetentionRulesRetentionTypeEnum() PolicyMonthlyRetentionRulesRetentionTypeEnum {
	return PolicyMonthlyRetentionRulesRetentionTypeEnum{
		WEEKLY: PolicyMonthlyRetentionRulesRetentionType{
			value: "WEEKLY",
		},
		MONTHLY: PolicyMonthlyRetentionRulesRetentionType{
			value: "MONTHLY",
		},
	}
}

func (c PolicyMonthlyRetentionRulesRetentionType) Value() string {
	return c.value
}

func (c PolicyMonthlyRetentionRulesRetentionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyMonthlyRetentionRulesRetentionType) UnmarshalJSON(b []byte) error {
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

type PolicyMonthlyRetentionRulesRetentionWeeks struct {
	value string
}

type PolicyMonthlyRetentionRulesRetentionWeeksEnum struct {
	FIRST  PolicyMonthlyRetentionRulesRetentionWeeks
	SECOND PolicyMonthlyRetentionRulesRetentionWeeks
	THIRD  PolicyMonthlyRetentionRulesRetentionWeeks
	FOURTH PolicyMonthlyRetentionRulesRetentionWeeks
	LAST   PolicyMonthlyRetentionRulesRetentionWeeks
}

func GetPolicyMonthlyRetentionRulesRetentionWeeksEnum() PolicyMonthlyRetentionRulesRetentionWeeksEnum {
	return PolicyMonthlyRetentionRulesRetentionWeeksEnum{
		FIRST: PolicyMonthlyRetentionRulesRetentionWeeks{
			value: "FIRST",
		},
		SECOND: PolicyMonthlyRetentionRulesRetentionWeeks{
			value: "SECOND",
		},
		THIRD: PolicyMonthlyRetentionRulesRetentionWeeks{
			value: "THIRD",
		},
		FOURTH: PolicyMonthlyRetentionRulesRetentionWeeks{
			value: "FOURTH",
		},
		LAST: PolicyMonthlyRetentionRulesRetentionWeeks{
			value: "LAST",
		},
	}
}

func (c PolicyMonthlyRetentionRulesRetentionWeeks) Value() string {
	return c.value
}

func (c PolicyMonthlyRetentionRulesRetentionWeeks) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyMonthlyRetentionRulesRetentionWeeks) UnmarshalJSON(b []byte) error {
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

type PolicyMonthlyRetentionRulesDaysOfWeek struct {
	value string
}

type PolicyMonthlyRetentionRulesDaysOfWeekEnum struct {
	MO PolicyMonthlyRetentionRulesDaysOfWeek
	TU PolicyMonthlyRetentionRulesDaysOfWeek
	WE PolicyMonthlyRetentionRulesDaysOfWeek
	TH PolicyMonthlyRetentionRulesDaysOfWeek
	FR PolicyMonthlyRetentionRulesDaysOfWeek
	SA PolicyMonthlyRetentionRulesDaysOfWeek
	SU PolicyMonthlyRetentionRulesDaysOfWeek
}

func GetPolicyMonthlyRetentionRulesDaysOfWeekEnum() PolicyMonthlyRetentionRulesDaysOfWeekEnum {
	return PolicyMonthlyRetentionRulesDaysOfWeekEnum{
		MO: PolicyMonthlyRetentionRulesDaysOfWeek{
			value: "MO",
		},
		TU: PolicyMonthlyRetentionRulesDaysOfWeek{
			value: "TU",
		},
		WE: PolicyMonthlyRetentionRulesDaysOfWeek{
			value: "WE",
		},
		TH: PolicyMonthlyRetentionRulesDaysOfWeek{
			value: "TH",
		},
		FR: PolicyMonthlyRetentionRulesDaysOfWeek{
			value: "FR",
		},
		SA: PolicyMonthlyRetentionRulesDaysOfWeek{
			value: "SA",
		},
		SU: PolicyMonthlyRetentionRulesDaysOfWeek{
			value: "SU",
		},
	}
}

func (c PolicyMonthlyRetentionRulesDaysOfWeek) Value() string {
	return c.value
}

func (c PolicyMonthlyRetentionRulesDaysOfWeek) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyMonthlyRetentionRulesDaysOfWeek) UnmarshalJSON(b []byte) error {
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
