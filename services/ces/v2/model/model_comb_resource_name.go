package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CombResourceName 资源名称匹配规则
type CombResourceName struct {

	// **参数解释** 资源名称条件值 **约束限制** 不涉及 **取值范围** 长度[1,128]个字符 **默认取值** 不涉及
	ResourceName *string `json:"resource_name,omitempty"`

	// **参数解释** 实例操作符，含义是真实资源的名称与资源名称条件值的运算关系。 **约束限制** 不涉及 **取值范围** - include: 表示包含 - prefix: 表示前缀 - suffix: 表示后缀 - notInclude: 表示不包含 - equal: 表示相等 **默认取值** 不涉及
	Operator CombResourceNameOperator `json:"operator"`

	// **参数解释** 资源名称忽略大小写 **约束限制** 不涉及 **取值范围** - true: 名称忽略大小写 - false: 名称不忽略大小写 **默认取值** false
	ResourceNameIsIgnoreCase *bool `json:"resource_name_is_ignore_case,omitempty"`
}

func (o CombResourceName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CombResourceName struct{}"
	}

	return strings.Join([]string{"CombResourceName", string(data)}, " ")
}

type CombResourceNameOperator struct {
	value string
}

type CombResourceNameOperatorEnum struct {
	INCLUDE     CombResourceNameOperator
	PREFIX      CombResourceNameOperator
	SUFFIX      CombResourceNameOperator
	NOT_INCLUDE CombResourceNameOperator
	EQUAL       CombResourceNameOperator
}

func GetCombResourceNameOperatorEnum() CombResourceNameOperatorEnum {
	return CombResourceNameOperatorEnum{
		INCLUDE: CombResourceNameOperator{
			value: "include",
		},
		PREFIX: CombResourceNameOperator{
			value: "prefix",
		},
		SUFFIX: CombResourceNameOperator{
			value: "suffix",
		},
		NOT_INCLUDE: CombResourceNameOperator{
			value: "notInclude",
		},
		EQUAL: CombResourceNameOperator{
			value: "equal",
		},
	}
}

func (c CombResourceNameOperator) Value() string {
	return c.value
}

func (c CombResourceNameOperator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CombResourceNameOperator) UnmarshalJSON(b []byte) error {
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
