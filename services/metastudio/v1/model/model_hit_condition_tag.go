package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HitConditionTag 命中条件定义
type HitConditionTag struct {

	// **参数解释**： 事件内容关键字段 > * event_type为1,2,3,4：与LiveEventReport中event.content中反序列化后的JSON字段对应。如：弹幕事件上报事件。   {     \"timestamp\": 1694481224245,     \"type\": 1,     \"content\": \"{\\\"user\\\":{\\\"userId\\\":\\\"2027271526\\\",\\\"name\\\":\\\"***\\\",\\\"level\\\":17,\\\"badge\\\":\\\"\\\",\\\"badgeLevel\\\":0},\\\"content\\\":\\\"***\\\"}\"   }   匹配弹幕内容，填写content；匹配用户平台等级，填写level。 > * 10：固定填写content即可。  **约束限制**： 不涉及 **取值范围**： 字符长度0-256位 **默认取值**： 不涉及
	Tag *string `json:"tag,omitempty"`

	// **参数解释**： 字段取值处理 **约束限制**： 不涉及 **取值范围**： * SUM：累计 * AVG：平均 * COUNT：计数 * NONE：无处理
	Operation *HitConditionTagOperation `json:"operation,omitempty"`

	// **参数解释**： 匹配类型。关键词匹配建议使用REGEX。 **约束限制**： 不涉及。 **取值范围**： * EQUAL: 完全相等 * REGEX：正则匹配 * MATH_GT：数值大于 * MATH_GE：数值大于等于  * MATH_LT：数值小于 * MATH_LE：数值小于等于 * MATH_EQ：数值相等  **默认取值**： 不涉及
	Match *HitConditionTagMatch `json:"match,omitempty"`

	// **参数解释**： 匹配值。 **约束限制**： 不涉及 **取值范围**： 字符长度0-1024 **默认取值**： 不涉及。
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
