package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BackendConstant struct {

	// 常量参数名
	Name *string `json:"name,omitempty"`

	// 常量参数类型
	Type *BackendConstantType `json:"type,omitempty"`

	// 常量参数位置
	Position *BackendConstantPosition `json:"position,omitempty"`

	// 常量参数描述
	Description *string `json:"description,omitempty"`

	// 常量参数值
	Value *string `json:"value,omitempty"`
}

func (o BackendConstant) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackendConstant struct{}"
	}

	return strings.Join([]string{"BackendConstant", string(data)}, " ")
}

type BackendConstantType struct {
	value string
}

type BackendConstantTypeEnum struct {
	REQUEST_PARAMETER_TYPE_NUMBER BackendConstantType
	REQUEST_PARAMETER_TYPE_STRING BackendConstantType
}

func GetBackendConstantTypeEnum() BackendConstantTypeEnum {
	return BackendConstantTypeEnum{
		REQUEST_PARAMETER_TYPE_NUMBER: BackendConstantType{
			value: "REQUEST_PARAMETER_TYPE_NUMBER",
		},
		REQUEST_PARAMETER_TYPE_STRING: BackendConstantType{
			value: "REQUEST_PARAMETER_TYPE_STRING",
		},
	}
}

func (c BackendConstantType) Value() string {
	return c.value
}

func (c BackendConstantType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendConstantType) UnmarshalJSON(b []byte) error {
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

type BackendConstantPosition struct {
	value string
}

type BackendConstantPositionEnum struct {
	REQUEST_PARAMETER_POSITION_PATH   BackendConstantPosition
	REQUEST_PARAMETER_POSITION_HEADER BackendConstantPosition
	REQUEST_PARAMETER_POSITION_QUERY  BackendConstantPosition
}

func GetBackendConstantPositionEnum() BackendConstantPositionEnum {
	return BackendConstantPositionEnum{
		REQUEST_PARAMETER_POSITION_PATH: BackendConstantPosition{
			value: "REQUEST_PARAMETER_POSITION_PATH",
		},
		REQUEST_PARAMETER_POSITION_HEADER: BackendConstantPosition{
			value: "REQUEST_PARAMETER_POSITION_HEADER",
		},
		REQUEST_PARAMETER_POSITION_QUERY: BackendConstantPosition{
			value: "REQUEST_PARAMETER_POSITION_QUERY",
		},
	}
}

func (c BackendConstantPosition) Value() string {
	return c.value
}

func (c BackendConstantPosition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendConstantPosition) UnmarshalJSON(b []byte) error {
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
