package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PolicyWeeklyRetentionRules struct {

	// 设置每个星期中的指定天为周备备份
	DaysOfWeek *[]PolicyWeeklyRetentionRulesDaysOfWeek `json:"days_of_week,omitempty"`

	// 周备的保留时间，取值范围为1-5200，以及-1，单位为周，-1代表周备策略不启用
	RetentionDurationPeriods *int32 `json:"retention_duration_periods,omitempty"`
}

func (o PolicyWeeklyRetentionRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyWeeklyRetentionRules struct{}"
	}

	return strings.Join([]string{"PolicyWeeklyRetentionRules", string(data)}, " ")
}

type PolicyWeeklyRetentionRulesDaysOfWeek struct {
	value string
}

type PolicyWeeklyRetentionRulesDaysOfWeekEnum struct {
	MO PolicyWeeklyRetentionRulesDaysOfWeek
	TU PolicyWeeklyRetentionRulesDaysOfWeek
	WE PolicyWeeklyRetentionRulesDaysOfWeek
	TH PolicyWeeklyRetentionRulesDaysOfWeek
	FR PolicyWeeklyRetentionRulesDaysOfWeek
	SA PolicyWeeklyRetentionRulesDaysOfWeek
	SU PolicyWeeklyRetentionRulesDaysOfWeek
}

func GetPolicyWeeklyRetentionRulesDaysOfWeekEnum() PolicyWeeklyRetentionRulesDaysOfWeekEnum {
	return PolicyWeeklyRetentionRulesDaysOfWeekEnum{
		MO: PolicyWeeklyRetentionRulesDaysOfWeek{
			value: "MO",
		},
		TU: PolicyWeeklyRetentionRulesDaysOfWeek{
			value: "TU",
		},
		WE: PolicyWeeklyRetentionRulesDaysOfWeek{
			value: "WE",
		},
		TH: PolicyWeeklyRetentionRulesDaysOfWeek{
			value: "TH",
		},
		FR: PolicyWeeklyRetentionRulesDaysOfWeek{
			value: "FR",
		},
		SA: PolicyWeeklyRetentionRulesDaysOfWeek{
			value: "SA",
		},
		SU: PolicyWeeklyRetentionRulesDaysOfWeek{
			value: "SU",
		},
	}
}

func (c PolicyWeeklyRetentionRulesDaysOfWeek) Value() string {
	return c.value
}

func (c PolicyWeeklyRetentionRulesDaysOfWeek) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyWeeklyRetentionRulesDaysOfWeek) UnmarshalJSON(b []byte) error {
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
