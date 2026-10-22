package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AbuseCondition 算力滥用阈值条件
type AbuseCondition struct {

	// 算力消耗统计维度。input: 仅输入, output: 仅输出, input_output: 输入+输出。call_count仅支持input
	Dimension *AbuseConditionDimension `json:"dimension,omitempty"`

	// 阈值指标类型。call_count: 调用次数, char_count: 字符数
	Metric *AbuseConditionMetric `json:"metric,omitempty"`

	// 阈值数值，必须大于0；call_count最大1000000，char_count最大100000000000000
	Value *int64 `json:"value,omitempty"`
}

func (o AbuseCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AbuseCondition struct{}"
	}

	return strings.Join([]string{"AbuseCondition", string(data)}, " ")
}

type AbuseConditionDimension struct {
	value string
}

type AbuseConditionDimensionEnum struct {
	INPUT        AbuseConditionDimension
	OUTPUT       AbuseConditionDimension
	INPUT_OUTPUT AbuseConditionDimension
}

func GetAbuseConditionDimensionEnum() AbuseConditionDimensionEnum {
	return AbuseConditionDimensionEnum{
		INPUT: AbuseConditionDimension{
			value: "input",
		},
		OUTPUT: AbuseConditionDimension{
			value: "output",
		},
		INPUT_OUTPUT: AbuseConditionDimension{
			value: "input_output",
		},
	}
}

func (c AbuseConditionDimension) Value() string {
	return c.value
}

func (c AbuseConditionDimension) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AbuseConditionDimension) UnmarshalJSON(b []byte) error {
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

type AbuseConditionMetric struct {
	value string
}

type AbuseConditionMetricEnum struct {
	CALL_COUNT AbuseConditionMetric
	CHAR_COUNT AbuseConditionMetric
}

func GetAbuseConditionMetricEnum() AbuseConditionMetricEnum {
	return AbuseConditionMetricEnum{
		CALL_COUNT: AbuseConditionMetric{
			value: "call_count",
		},
		CHAR_COUNT: AbuseConditionMetric{
			value: "char_count",
		},
	}
}

func (c AbuseConditionMetric) Value() string {
	return c.value
}

func (c AbuseConditionMetric) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AbuseConditionMetric) UnmarshalJSON(b []byte) error {
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
