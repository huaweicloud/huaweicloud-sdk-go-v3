package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobInputResp 作业输入。
type JobInputResp struct {

	// **参数解释**：输入数据的名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：输入项类型。 **取值范围**：枚举值如下： - dataset：数据集 - obs：OBS - data_selector：数据选择
	Type *JobInputRespType `json:"type,omitempty"`

	// **参数解释**：输入项数据。
	Data *interface{} `json:"data,omitempty"`

	// **参数解释**：输入项的值。
	Value *interface{} `json:"value,omitempty"`
}

func (o JobInputResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInputResp struct{}"
	}

	return strings.Join([]string{"JobInputResp", string(data)}, " ")
}

type JobInputRespType struct {
	value string
}

type JobInputRespTypeEnum struct {
	DATASET JobInputRespType
	OBSOBS  JobInputRespType
}

func GetJobInputRespTypeEnum() JobInputRespTypeEnum {
	return JobInputRespTypeEnum{
		DATASET: JobInputRespType{
			value: "dataset：数据集。",
		},
		OBSOBS: JobInputRespType{
			value: "obs：OBS文件。",
		},
	}
}

func (c JobInputRespType) Value() string {
	return c.value
}

func (c JobInputRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInputRespType) UnmarshalJSON(b []byte) error {
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
