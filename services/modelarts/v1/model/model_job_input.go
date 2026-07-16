package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobInput 作业输入。
type JobInput struct {

	// 输入数据的名称，支持1到64位只包含英文、数字、下划线（_）和中划线（-）的字符。
	Name *string `json:"name,omitempty"`

	// 输入项类型。枚举值如下： - dataset：数据集 - obs：OBS - data_selector：数据选择
	Type *JobInputType `json:"type,omitempty"`

	// 输入项数据。
	Data *interface{} `json:"data,omitempty"`

	// 输入项的值。
	Value *interface{} `json:"value,omitempty"`
}

func (o JobInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInput struct{}"
	}

	return strings.Join([]string{"JobInput", string(data)}, " ")
}

type JobInputType struct {
	value string
}

type JobInputTypeEnum struct {
	DATASET JobInputType
	OBSOBS  JobInputType
}

func GetJobInputTypeEnum() JobInputTypeEnum {
	return JobInputTypeEnum{
		DATASET: JobInputType{
			value: "dataset：数据集。",
		},
		OBSOBS: JobInputType{
			value: "obs：OBS文件。",
		},
	}
}

func (c JobInputType) Value() string {
	return c.value
}

func (c JobInputType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInputType) UnmarshalJSON(b []byte) error {
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
