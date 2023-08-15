package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CustomPropsTaskData 自定义属性任务的请求体
type CustomPropsTaskData struct {

	// 自定义属性名称
	Name string `json:"name"`

	// 自定义属性描述信息
	Description *string `json:"description,omitempty"`

	// 属性预测类型
	Type CustomPropsTaskDataType `json:"type"`

	// 用于建模的smiles列表
	Smiles []string `json:"smiles"`

	// 用于建模的属性值列表
	Values []float32 `json:"values"`
}

func (o CustomPropsTaskData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomPropsTaskData struct{}"
	}

	return strings.Join([]string{"CustomPropsTaskData", string(data)}, " ")
}

type CustomPropsTaskDataType struct {
	value string
}

type CustomPropsTaskDataTypeEnum struct {
	BINARY    CustomPropsTaskDataType
	NUMERICAL CustomPropsTaskDataType
}

func GetCustomPropsTaskDataTypeEnum() CustomPropsTaskDataTypeEnum {
	return CustomPropsTaskDataTypeEnum{
		BINARY: CustomPropsTaskDataType{
			value: "binary",
		},
		NUMERICAL: CustomPropsTaskDataType{
			value: "numerical",
		},
	}
}

func (c CustomPropsTaskDataType) Value() string {
	return c.value
}

func (c CustomPropsTaskDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomPropsTaskDataType) UnmarshalJSON(b []byte) error {
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
