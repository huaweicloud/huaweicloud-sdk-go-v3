package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Data 数据。
type Data struct {

	// 训练数据的名称。填写1-64位，仅包含英文、数字、下划线（_）和中划线（-），并且以英文开头的名称。
	Name *string `json:"name,omitempty"`

	// 数据来源的类型，可选值为dataset、obs、swr、model、label_task、service、image。
	Type *DataType `json:"type,omitempty"`

	// 数据的值。
	Value map[string]interface{} `json:"value,omitempty"`

	// 使用数据的节点。
	UsedSteps *[]string `json:"used_steps,omitempty"`
}

func (o Data) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Data struct{}"
	}

	return strings.Join([]string{"Data", string(data)}, " ")
}

type DataType struct {
	value string
}

type DataTypeEnum struct {
	DATASET DataType
	OBSOBS  DataType
}

func GetDataTypeEnum() DataTypeEnum {
	return DataTypeEnum{
		DATASET: DataType{
			value: "dataset：数据集。",
		},
		OBSOBS: DataType{
			value: "obs：OBS文件。",
		},
	}
}

func (c DataType) Value() string {
	return c.value
}

func (c DataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataType) UnmarshalJSON(b []byte) error {
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
