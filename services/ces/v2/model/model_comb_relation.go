package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CombRelation 组合匹配参数
type CombRelation struct {

	// 逻辑运算符  ALL 所有条件匹配成功  ANY 任意条件匹配成功
	LogicalOperator CombRelationLogicalOperator `json:"logical_operator"`

	// 组合匹配资源分组的匹配条件
	Conditions []Condition `json:"conditions"`
}

func (o CombRelation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CombRelation struct{}"
	}

	return strings.Join([]string{"CombRelation", string(data)}, " ")
}

type CombRelationLogicalOperator struct {
	value string
}

type CombRelationLogicalOperatorEnum struct {
	ALL CombRelationLogicalOperator
	ANY CombRelationLogicalOperator
}

func GetCombRelationLogicalOperatorEnum() CombRelationLogicalOperatorEnum {
	return CombRelationLogicalOperatorEnum{
		ALL: CombRelationLogicalOperator{
			value: "ALL",
		},
		ANY: CombRelationLogicalOperator{
			value: "ANY",
		},
	}
}

func (c CombRelationLogicalOperator) Value() string {
	return c.value
}

func (c CombRelationLogicalOperator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CombRelationLogicalOperator) UnmarshalJSON(b []byte) error {
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
