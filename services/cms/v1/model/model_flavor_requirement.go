package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 对规格的资源需求约束
type FlavorRequirement struct {
	VcpuCount *IntegerRange `json:"vcpu_count,omitempty"`

	MemoryMb *IntegerRange `json:"memory_mb,omitempty"`

	// 可选CPU制造商，不填表示接受所有
	CpuManufacturers *[]FlavorRequirementCpuManufacturers `json:"cpu_manufacturers,omitempty"`

	MemoryGbPerVcpu *DoubleRange `json:"memory_gb_per_vcpu,omitempty"`

	// 接受的资源代系，不填表示接受所有
	InstanceGenerations *[]FlavorRequirementInstanceGenerations `json:"instance_generations,omitempty"`
}

func (o FlavorRequirement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorRequirement struct{}"
	}

	return strings.Join([]string{"FlavorRequirement", string(data)}, " ")
}

type FlavorRequirementCpuManufacturers struct {
	value string
}

type FlavorRequirementCpuManufacturersEnum struct {
	INTEL FlavorRequirementCpuManufacturers
	AMD   FlavorRequirementCpuManufacturers
	OTHER FlavorRequirementCpuManufacturers
}

func GetFlavorRequirementCpuManufacturersEnum() FlavorRequirementCpuManufacturersEnum {
	return FlavorRequirementCpuManufacturersEnum{
		INTEL: FlavorRequirementCpuManufacturers{
			value: "INTEL",
		},
		AMD: FlavorRequirementCpuManufacturers{
			value: "AMD",
		},
		OTHER: FlavorRequirementCpuManufacturers{
			value: "OTHER",
		},
	}
}

func (c FlavorRequirementCpuManufacturers) Value() string {
	return c.value
}

func (c FlavorRequirementCpuManufacturers) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlavorRequirementCpuManufacturers) UnmarshalJSON(b []byte) error {
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

type FlavorRequirementInstanceGenerations struct {
	value string
}

type FlavorRequirementInstanceGenerationsEnum struct {
	CURRENT  FlavorRequirementInstanceGenerations
	PREVIOUS FlavorRequirementInstanceGenerations
}

func GetFlavorRequirementInstanceGenerationsEnum() FlavorRequirementInstanceGenerationsEnum {
	return FlavorRequirementInstanceGenerationsEnum{
		CURRENT: FlavorRequirementInstanceGenerations{
			value: "CURRENT",
		},
		PREVIOUS: FlavorRequirementInstanceGenerations{
			value: "PREVIOUS",
		},
	}
}

func (c FlavorRequirementInstanceGenerations) Value() string {
	return c.value
}

func (c FlavorRequirementInstanceGenerations) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlavorRequirementInstanceGenerations) UnmarshalJSON(b []byte) error {
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
