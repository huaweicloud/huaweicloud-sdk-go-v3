package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AlarmRuleParam 阈值规则实体
type AlarmRuleParam struct {

	// 是否启用通知。
	ActionEnabled *bool `json:"action_enabled,omitempty"`

	// 告警状态通知列表。
	AlarmActions *[]string `json:"alarm_actions,omitempty"`

	// 告警清除建议。
	AlarmAdvice *string `json:"alarm_advice,omitempty"`

	// 阈值规则描述。
	AlarmDescription *string `json:"alarm_description,omitempty"`

	// 告警级别。1：紧急，2：重要，3：一般，4：提示。
	AlarmLevel AlarmRuleParamAlarmLevel `json:"alarm_level"`

	// 阈值规则名称。
	AlarmRuleName string `json:"alarm_rule_name"`

	// 超限条件。
	ComparisonOperator string `json:"comparison_operator"`

	// 时间序列维度。
	Dimensions []Dimension `json:"dimensions"`

	// 间隔周期。
	EvaluationPeriods int32 `json:"evaluation_periods"`

	// 阈值规则是否启用。
	IdTurnOn *bool `json:"id_turn_on,omitempty"`

	// 数据不足通知列表。
	InsufficientDataActions *[]string `json:"insufficient_data_actions,omitempty"`

	// 时间序列名称。名称长度取值范围为1~255个字符。
	MetricName string `json:"metric_name"`

	// 时间序列命名空间。
	Namespace string `json:"namespace"`

	// 正常状态通知列表。
	OkActions *[]string `json:"ok_actions,omitempty"`

	// 统计周期。
	Period int32 `json:"period"`

	// 统计方式。
	Statistic AlarmRuleParamStatistic `json:"statistic"`

	// 超限值。
	Threshold string `json:"threshold"`

	// 时间序列单位
	Unit string `json:"unit"`
}

func (o AlarmRuleParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmRuleParam struct{}"
	}

	return strings.Join([]string{"AlarmRuleParam", string(data)}, " ")
}

type AlarmRuleParamAlarmLevel struct {
	value int32
}

type AlarmRuleParamAlarmLevelEnum struct {
	E_1 AlarmRuleParamAlarmLevel
	E_2 AlarmRuleParamAlarmLevel
	E_3 AlarmRuleParamAlarmLevel
	E_4 AlarmRuleParamAlarmLevel
}

func GetAlarmRuleParamAlarmLevelEnum() AlarmRuleParamAlarmLevelEnum {
	return AlarmRuleParamAlarmLevelEnum{
		E_1: AlarmRuleParamAlarmLevel{
			value: 1,
		}, E_2: AlarmRuleParamAlarmLevel{
			value: 2,
		}, E_3: AlarmRuleParamAlarmLevel{
			value: 3,
		}, E_4: AlarmRuleParamAlarmLevel{
			value: 4,
		},
	}
}

func (c AlarmRuleParamAlarmLevel) Value() int32 {
	return c.value
}

func (c AlarmRuleParamAlarmLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmRuleParamAlarmLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type AlarmRuleParamStatistic struct {
	value string
}

type AlarmRuleParamStatisticEnum struct {
	MAXIMUM      AlarmRuleParamStatistic
	MINIMUM      AlarmRuleParamStatistic
	AVERAGE      AlarmRuleParamStatistic
	SUM          AlarmRuleParamStatistic
	SAMPLE_COUNT AlarmRuleParamStatistic
}

func GetAlarmRuleParamStatisticEnum() AlarmRuleParamStatisticEnum {
	return AlarmRuleParamStatisticEnum{
		MAXIMUM: AlarmRuleParamStatistic{
			value: "maximum",
		},
		MINIMUM: AlarmRuleParamStatistic{
			value: "minimum",
		},
		AVERAGE: AlarmRuleParamStatistic{
			value: "average",
		},
		SUM: AlarmRuleParamStatistic{
			value: "sum",
		},
		SAMPLE_COUNT: AlarmRuleParamStatistic{
			value: "sampleCount",
		},
	}
}

func (c AlarmRuleParamStatistic) Value() string {
	return c.value
}

func (c AlarmRuleParamStatistic) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmRuleParamStatistic) UnmarshalJSON(b []byte) error {
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
