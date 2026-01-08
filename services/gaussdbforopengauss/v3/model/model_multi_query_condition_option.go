package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MultiQueryConditionOption struct {

	// **参数解释**： 查询字段名称。 **约束限制**： 只支持字符串\"query\"。 **取值范围**： 由英文字母（大小写）、数字或下划线组成，长度为 1 至 128 个字符。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**: 组合条件类型。 **约束限制**: 不涉及。 **取值范围**: 仅限字符串：\"AND\"、\"OR\"。 **默认取值**: 不涉及。
	Condition MultiQueryConditionOptionCondition `json:"condition"`

	// **参数解释**: 多个过滤检索条件内容集合。 **约束限制**: 只支持为true进行模糊查询。 **取值范围**: - true：模糊查询。 - false：精确匹配。  **默认取值**: true
	IsFuzzy *bool `json:"is_fuzzy,omitempty"`

	// **参数解释**: 多个过滤检索条件内容集合。由 1 至 5 个字符串组成的列表。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	Values []string `json:"values"`
}

func (o MultiQueryConditionOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiQueryConditionOption struct{}"
	}

	return strings.Join([]string{"MultiQueryConditionOption", string(data)}, " ")
}

type MultiQueryConditionOptionCondition struct {
	value string
}

type MultiQueryConditionOptionConditionEnum struct {
	OR  MultiQueryConditionOptionCondition
	AND MultiQueryConditionOptionCondition
}

func GetMultiQueryConditionOptionConditionEnum() MultiQueryConditionOptionConditionEnum {
	return MultiQueryConditionOptionConditionEnum{
		OR: MultiQueryConditionOptionCondition{
			value: "OR",
		},
		AND: MultiQueryConditionOptionCondition{
			value: "AND",
		},
	}
}

func (c MultiQueryConditionOptionCondition) Value() string {
	return c.value
}

func (c MultiQueryConditionOptionCondition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MultiQueryConditionOptionCondition) UnmarshalJSON(b []byte) error {
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
