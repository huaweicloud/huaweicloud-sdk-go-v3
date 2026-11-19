package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ComputeFlavorGroupDetail struct {

	// **参数解释**: 性能规格。 **取值范围**: - normal：独享型（1:8）。 - normal2：独享型（1:4）。 - armFlavors：鲲鹏独享型（1:8）。 - exclusive：独享型（1:4），仅基础版集中式实例支持。 - armExclusive：鲲鹏独享型（1:4），仅基础版集中式实例支持。 - economical：惠选型（1:4）。 - economical2：惠选型（1:8）。 - armFlavors2：鲲鹏独享型（1:4）。 - general：通用型（1:4）。
	GroupType *ComputeFlavorGroupDetailGroupType `json:"group_type,omitempty"`

	// **参数解释**: 计算规格列表。 **取值范围**: 不涉及。
	ComputeFlavors *[]ComputeFlavorItemDetail `json:"compute_flavors,omitempty"`
}

func (o ComputeFlavorGroupDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputeFlavorGroupDetail struct{}"
	}

	return strings.Join([]string{"ComputeFlavorGroupDetail", string(data)}, " ")
}

type ComputeFlavorGroupDetailGroupType struct {
	value string
}

type ComputeFlavorGroupDetailGroupTypeEnum struct {
	NORMAL        ComputeFlavorGroupDetailGroupType
	NORMAL2       ComputeFlavorGroupDetailGroupType
	ARM_FLAVORS   ComputeFlavorGroupDetailGroupType
	EXCLUSIVE     ComputeFlavorGroupDetailGroupType
	ARM_EXCLUSIVE ComputeFlavorGroupDetailGroupType
	ECONOMICAL    ComputeFlavorGroupDetailGroupType
	ECONOMICAL2   ComputeFlavorGroupDetailGroupType
	ARM_FLAVORS2  ComputeFlavorGroupDetailGroupType
	GENERAL       ComputeFlavorGroupDetailGroupType
}

func GetComputeFlavorGroupDetailGroupTypeEnum() ComputeFlavorGroupDetailGroupTypeEnum {
	return ComputeFlavorGroupDetailGroupTypeEnum{
		NORMAL: ComputeFlavorGroupDetailGroupType{
			value: "normal",
		},
		NORMAL2: ComputeFlavorGroupDetailGroupType{
			value: "normal2",
		},
		ARM_FLAVORS: ComputeFlavorGroupDetailGroupType{
			value: "armFlavors",
		},
		EXCLUSIVE: ComputeFlavorGroupDetailGroupType{
			value: "exclusive",
		},
		ARM_EXCLUSIVE: ComputeFlavorGroupDetailGroupType{
			value: "armExclusive",
		},
		ECONOMICAL: ComputeFlavorGroupDetailGroupType{
			value: "economical",
		},
		ECONOMICAL2: ComputeFlavorGroupDetailGroupType{
			value: "economical2",
		},
		ARM_FLAVORS2: ComputeFlavorGroupDetailGroupType{
			value: "armFlavors2",
		},
		GENERAL: ComputeFlavorGroupDetailGroupType{
			value: "general",
		},
	}
}

func (c ComputeFlavorGroupDetailGroupType) Value() string {
	return c.value
}

func (c ComputeFlavorGroupDetailGroupType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComputeFlavorGroupDetailGroupType) UnmarshalJSON(b []byte) error {
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
