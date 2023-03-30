package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 分子约束
type MoleculeConstraint struct {

	// 属性名称
	Name *string `json:"name,omitempty"`

	// 属性约束类型
	Type MoleculeConstraintType `json:"type"`

	// 属性约束类型bool的参数
	Bool *bool `json:"bool,omitempty"`

	// 属性约束类型range的参数
	Range *[]float32 `json:"range,omitempty"`

	Struct *StructureConstraintParams `json:"struct,omitempty"`

	// 属性约束类型minimize和maximize的参数
	Quantiles *[]float32 `json:"quantiles,omitempty"`
}

func (o MoleculeConstraint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoleculeConstraint struct{}"
	}

	return strings.Join([]string{"MoleculeConstraint", string(data)}, " ")
}

type MoleculeConstraintType struct {
	value string
}

type MoleculeConstraintTypeEnum struct {
	BOOL     MoleculeConstraintType
	RANGE    MoleculeConstraintType
	STRUCT   MoleculeConstraintType
	MINIMIZE MoleculeConstraintType
	MAXIMIZE MoleculeConstraintType
}

func GetMoleculeConstraintTypeEnum() MoleculeConstraintTypeEnum {
	return MoleculeConstraintTypeEnum{
		BOOL: MoleculeConstraintType{
			value: "bool",
		},
		RANGE: MoleculeConstraintType{
			value: "range",
		},
		STRUCT: MoleculeConstraintType{
			value: "struct",
		},
		MINIMIZE: MoleculeConstraintType{
			value: "minimize",
		},
		MAXIMIZE: MoleculeConstraintType{
			value: "maximize",
		},
	}
}

func (c MoleculeConstraintType) Value() string {
	return c.value
}

func (c MoleculeConstraintType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MoleculeConstraintType) UnmarshalJSON(b []byte) error {
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
