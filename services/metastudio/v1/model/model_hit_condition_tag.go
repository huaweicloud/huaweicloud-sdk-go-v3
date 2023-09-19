package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HitConditionTag 命中条件定义
type HitConditionTag struct {

	// 事件内容关键字段
	Tag *string `json:"tag,omitempty"`

	// 字段处理 - SUM: 累计 - AVG：平均 - COUNT: 计数 - NONE: 无处理
	Operation *HitConditionTagOperation `json:"operation,omitempty"`

	// 匹配类型 - EQUAL: 完全相等 - REGEX：正则匹配 - MATH_GT：数值大于 - MATH_GE： 数值大于等于 - MATH_LT：数值小于 - MATH_LE：数值小于等于 - MATH_EQ：数值相等
	Match *HitConditionTagMatch `json:"match,omitempty"`

	// 匹配值
	Value *string `json:"value,omitempty"`
}

func (o HitConditionTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HitConditionTag struct{}"
	}

	return strings.Join([]string{"HitConditionTag", string(data)}, " ")
}

type HitConditionTagOperation struct {
	value string
}

type HitConditionTagOperationEnum struct {
	SUM   HitConditionTagOperation
	AVG   HitConditionTagOperation
	COUNT HitConditionTagOperation
	NONE  HitConditionTagOperation
}

func GetHitConditionTagOperationEnum() HitConditionTagOperationEnum {
	return HitConditionTagOperationEnum{
		SUM: HitConditionTagOperation{
			value: "SUM",
		},
		AVG: HitConditionTagOperation{
			value: "AVG",
		},
		COUNT: HitConditionTagOperation{
			value: "COUNT",
		},
		NONE: HitConditionTagOperation{
			value: "NONE",
		},
	}
}

func (c HitConditionTagOperation) Value() string {
	return c.value
}

func (c HitConditionTagOperation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HitConditionTagOperation) UnmarshalJSON(b []byte) error {
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

type HitConditionTagMatch struct {
	value string
}

type HitConditionTagMatchEnum struct {
	EQUAL   HitConditionTagMatch
	REGEX   HitConditionTagMatch
	MATH_GT HitConditionTagMatch
	MATH_GE HitConditionTagMatch
	MATH_LT HitConditionTagMatch
	MATH_LE HitConditionTagMatch
	MATH_EQ HitConditionTagMatch
}

func GetHitConditionTagMatchEnum() HitConditionTagMatchEnum {
	return HitConditionTagMatchEnum{
		EQUAL: HitConditionTagMatch{
			value: "EQUAL",
		},
		REGEX: HitConditionTagMatch{
			value: "REGEX",
		},
		MATH_GT: HitConditionTagMatch{
			value: "MATH_GT",
		},
		MATH_GE: HitConditionTagMatch{
			value: "MATH_GE",
		},
		MATH_LT: HitConditionTagMatch{
			value: "MATH_LT",
		},
		MATH_LE: HitConditionTagMatch{
			value: "MATH_LE",
		},
		MATH_EQ: HitConditionTagMatch{
			value: "MATH_EQ",
		},
	}
}

func (c HitConditionTagMatch) Value() string {
	return c.value
}

func (c HitConditionTagMatch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HitConditionTagMatch) UnmarshalJSON(b []byte) error {
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
