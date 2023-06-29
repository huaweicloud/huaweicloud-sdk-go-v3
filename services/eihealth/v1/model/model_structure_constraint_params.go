package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StructureConstraintParams 结构约束参数
type StructureConstraintParams struct {

	// 子结构SMILES
	Structs []string `json:"structs"`

	// 是否排除子结构
	Exclusive bool `json:"exclusive"`

	// 多个子结构之间的逻辑关系
	Operator *StructureConstraintParamsOperator `json:"operator,omitempty"`
}

func (o StructureConstraintParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructureConstraintParams struct{}"
	}

	return strings.Join([]string{"StructureConstraintParams", string(data)}, " ")
}

type StructureConstraintParamsOperator struct {
	value string
}

type StructureConstraintParamsOperatorEnum struct {
	OR  StructureConstraintParamsOperator
	AND StructureConstraintParamsOperator
}

func GetStructureConstraintParamsOperatorEnum() StructureConstraintParamsOperatorEnum {
	return StructureConstraintParamsOperatorEnum{
		OR: StructureConstraintParamsOperator{
			value: "or",
		},
		AND: StructureConstraintParamsOperator{
			value: "and",
		},
	}
}

func (c StructureConstraintParamsOperator) Value() string {
	return c.value
}

func (c StructureConstraintParamsOperator) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StructureConstraintParamsOperator) UnmarshalJSON(b []byte) error {
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
