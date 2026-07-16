package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataRequirementResp 工作流需要的数据。
type DataRequirementResp struct {

	// **参数解释**：训练数据的名称。 **取值范围**：不涉及。
	Name string `json:"name"`

	// **参数解释**：数据来源类型。 **取值范围**：枚举值如下： - dataset：数据集 - obs：OBS - swr：SWR - model_list：AI应用列表 - label_task：标注任务 - service：在线服务
	Type DataRequirementRespType `json:"type"`

	// **参数解释**：数据约束条件。
	Conditions *[]ConstraintResp `json:"conditions,omitempty"`

	// **参数解释**：数据的值。
	Value map[string]interface{} `json:"value,omitempty"`

	// **参数解释**：使用了这条数据的工作流节点。
	UsedSteps *[]string `json:"used_steps,omitempty"`

	// **参数解释**：延时参数标记。 **取值范围**： - true：延时 - false：不延时
	Delay *bool `json:"delay,omitempty"`
}

func (o DataRequirementResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataRequirementResp struct{}"
	}

	return strings.Join([]string{"DataRequirementResp", string(data)}, " ")
}

type DataRequirementRespType struct {
	value string
}

type DataRequirementRespTypeEnum struct {
	DATASET DataRequirementRespType
	OBSOBS  DataRequirementRespType
}

func GetDataRequirementRespTypeEnum() DataRequirementRespTypeEnum {
	return DataRequirementRespTypeEnum{
		DATASET: DataRequirementRespType{
			value: "dataset：数据集",
		},
		OBSOBS: DataRequirementRespType{
			value: "obs：OBS文件",
		},
	}
}

func (c DataRequirementRespType) Value() string {
	return c.value
}

func (c DataRequirementRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataRequirementRespType) UnmarshalJSON(b []byte) error {
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
