package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataRequirement 工作流需要的数据。
type DataRequirement struct {

	// 训练数据的名称。填写1-64位，仅包含英文、数字、下划线（_）和中划线（-），并且以英文开头的名称。
	Name string `json:"name"`

	// 数据来源类型。枚举值如下： - dataset：数据集 - obs：OBS - swr：SWR - model_list：AI应用列表 - label_task：标注任务 - service：在线服务
	Type DataRequirementType `json:"type"`

	// 数据约束条件。
	Conditions *[]Constraint `json:"conditions,omitempty"`

	// 数据的值。
	Value map[string]interface{} `json:"value,omitempty"`

	// 使用了这条数据的工作流节点。
	UsedSteps *[]string `json:"used_steps,omitempty"`

	// 延时参数标记。
	Delay *bool `json:"delay,omitempty"`
}

func (o DataRequirement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataRequirement struct{}"
	}

	return strings.Join([]string{"DataRequirement", string(data)}, " ")
}

type DataRequirementType struct {
	value string
}

type DataRequirementTypeEnum struct {
	DATASET DataRequirementType
	OBSOBS  DataRequirementType
}

func GetDataRequirementTypeEnum() DataRequirementTypeEnum {
	return DataRequirementTypeEnum{
		DATASET: DataRequirementType{
			value: "dataset：数据集",
		},
		OBSOBS: DataRequirementType{
			value: "obs：OBS文件",
		},
	}
}

func (c DataRequirementType) Value() string {
	return c.value
}

func (c DataRequirementType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataRequirementType) UnmarshalJSON(b []byte) error {
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
