package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DataClassificationGroupCombineRuleDto struct {

	// 规则名称
	Name string `json:"name"`

	// 密级ID
	SecrecyLevelId string `json:"secrecy_level_id"`

	// 条件表达式
	CombineExpression string `json:"combine_expression"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 分类ID
	CategoryId *string `json:"category_id,omitempty"`

	// 使能状态。
	Enable *bool `json:"enable,omitempty"`

	// 规则方式, COMBINE
	Method DataClassificationGroupCombineRuleDtoMethod `json:"method"`

	// 条件单规则
	SingleExpressions []DataClassificationSingleRuleDto `json:"single_expressions"`
}

func (o DataClassificationGroupCombineRuleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataClassificationGroupCombineRuleDto struct{}"
	}

	return strings.Join([]string{"DataClassificationGroupCombineRuleDto", string(data)}, " ")
}

type DataClassificationGroupCombineRuleDtoMethod struct {
	value string
}

type DataClassificationGroupCombineRuleDtoMethodEnum struct {
	COMBINE DataClassificationGroupCombineRuleDtoMethod
}

func GetDataClassificationGroupCombineRuleDtoMethodEnum() DataClassificationGroupCombineRuleDtoMethodEnum {
	return DataClassificationGroupCombineRuleDtoMethodEnum{
		COMBINE: DataClassificationGroupCombineRuleDtoMethod{
			value: "COMBINE",
		},
	}
}

func (c DataClassificationGroupCombineRuleDtoMethod) Value() string {
	return c.value
}

func (c DataClassificationGroupCombineRuleDtoMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataClassificationGroupCombineRuleDtoMethod) UnmarshalJSON(b []byte) error {
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
