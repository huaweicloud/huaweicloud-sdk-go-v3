package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExecuteActionRequestBody struct {

	// API版本。
	ApiVersion ExecuteActionRequestBodyApiVersion `json:"api_version"`

	// 资源种类。
	Kind ExecuteActionRequestBodyKind `json:"kind"`

	Metadata *ExecuteActionRequestBodyMetadata `json:"metadata,omitempty"`

	Spec *ActionOnComponentSpec `json:"spec,omitempty"`
}

func (o ExecuteActionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteActionRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteActionRequestBody", string(data)}, " ")
}

type ExecuteActionRequestBodyApiVersion struct {
	value string
}

type ExecuteActionRequestBodyApiVersionEnum struct {
	V1 ExecuteActionRequestBodyApiVersion
}

func GetExecuteActionRequestBodyApiVersionEnum() ExecuteActionRequestBodyApiVersionEnum {
	return ExecuteActionRequestBodyApiVersionEnum{
		V1: ExecuteActionRequestBodyApiVersion{
			value: "v1",
		},
	}
}

func (c ExecuteActionRequestBodyApiVersion) Value() string {
	return c.value
}

func (c ExecuteActionRequestBodyApiVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteActionRequestBodyApiVersion) UnmarshalJSON(b []byte) error {
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

type ExecuteActionRequestBodyKind struct {
	value string
}

type ExecuteActionRequestBodyKindEnum struct {
	ACTION ExecuteActionRequestBodyKind
}

func GetExecuteActionRequestBodyKindEnum() ExecuteActionRequestBodyKindEnum {
	return ExecuteActionRequestBodyKindEnum{
		ACTION: ExecuteActionRequestBodyKind{
			value: "Action",
		},
	}
}

func (c ExecuteActionRequestBodyKind) Value() string {
	return c.value
}

func (c ExecuteActionRequestBodyKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteActionRequestBodyKind) UnmarshalJSON(b []byte) error {
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
