package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 订阅的事件目标转换规则
type SubscriptionTargetTransform struct {

	// 转换规则类型
	Type SubscriptionTargetTransformType `json:"type"`

	// - 转换类型规则为常量时，字段为常量内容定义；  - 转换类型规则为变量时，为变量定义，内容必须为JsonObject字符串。   - 变量最多支持100个，且不支持嵌套结构定义；   - 变量名由字母、数字、点、下划线和中划线组成，必须字母或数字开头不能以HC.开头，长度不超过64个字符；   - 变量值表达式支持常量或JsonPath表达式，字符串长度不超过1024个字符。
	Value *string `json:"value,omitempty"`

	// 转换规则类型为变量时，规则内容由模板定义，支持对已定义变量的引用。
	Template *string `json:"template,omitempty"`
}

func (o SubscriptionTargetTransform) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionTargetTransform struct{}"
	}

	return strings.Join([]string{"SubscriptionTargetTransform", string(data)}, " ")
}

type SubscriptionTargetTransformType struct {
	value string
}

type SubscriptionTargetTransformTypeEnum struct {
	ORIGINAL SubscriptionTargetTransformType
	CONSTANT SubscriptionTargetTransformType
	VARIABLE SubscriptionTargetTransformType
}

func GetSubscriptionTargetTransformTypeEnum() SubscriptionTargetTransformTypeEnum {
	return SubscriptionTargetTransformTypeEnum{
		ORIGINAL: SubscriptionTargetTransformType{
			value: "ORIGINAL",
		},
		CONSTANT: SubscriptionTargetTransformType{
			value: "CONSTANT",
		},
		VARIABLE: SubscriptionTargetTransformType{
			value: "VARIABLE",
		},
	}
}

func (c SubscriptionTargetTransformType) Value() string {
	return c.value
}

func (c SubscriptionTargetTransformType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubscriptionTargetTransformType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
