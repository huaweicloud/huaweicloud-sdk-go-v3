package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlavorConstraint 资源规格的筛选约束
type FlavorConstraint struct {

	// 接受的体系结构描述
	ArchitectureType *[]FlavorConstraintArchitectureType `json:"architecture_type,omitempty"`

	// 资源的需求约束
	FlavorRequirements *[]FlavorRequirement `json:"flavor_requirements,omitempty"`
}

func (o FlavorConstraint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorConstraint struct{}"
	}

	return strings.Join([]string{"FlavorConstraint", string(data)}, " ")
}

type FlavorConstraintArchitectureType struct {
	value string
}

type FlavorConstraintArchitectureTypeEnum struct {
	X86_64 FlavorConstraintArchitectureType
	ARM64  FlavorConstraintArchitectureType
}

func GetFlavorConstraintArchitectureTypeEnum() FlavorConstraintArchitectureTypeEnum {
	return FlavorConstraintArchitectureTypeEnum{
		X86_64: FlavorConstraintArchitectureType{
			value: "x86_64",
		},
		ARM64: FlavorConstraintArchitectureType{
			value: "arm64",
		},
	}
}

func (c FlavorConstraintArchitectureType) Value() string {
	return c.value
}

func (c FlavorConstraintArchitectureType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlavorConstraintArchitectureType) UnmarshalJSON(b []byte) error {
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
