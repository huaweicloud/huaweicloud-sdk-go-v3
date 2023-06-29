package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateComponentRequestBody struct {

	// API版本。
	ApiVersion CreateComponentRequestBodyApiVersion `json:"api_version"`

	// 资源种类。
	Kind CreateComponentRequestBodyKind `json:"kind"`

	Metadata *CreateComponentRequestBodyMetadata `json:"metadata,omitempty"`

	Spec *CreateComponentRequestBodySpec `json:"spec,omitempty"`
}

func (o CreateComponentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentRequestBody struct{}"
	}

	return strings.Join([]string{"CreateComponentRequestBody", string(data)}, " ")
}

type CreateComponentRequestBodyApiVersion struct {
	value string
}

type CreateComponentRequestBodyApiVersionEnum struct {
	V1 CreateComponentRequestBodyApiVersion
}

func GetCreateComponentRequestBodyApiVersionEnum() CreateComponentRequestBodyApiVersionEnum {
	return CreateComponentRequestBodyApiVersionEnum{
		V1: CreateComponentRequestBodyApiVersion{
			value: "v1",
		},
	}
}

func (c CreateComponentRequestBodyApiVersion) Value() string {
	return c.value
}

func (c CreateComponentRequestBodyApiVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateComponentRequestBodyApiVersion) UnmarshalJSON(b []byte) error {
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

type CreateComponentRequestBodyKind struct {
	value string
}

type CreateComponentRequestBodyKindEnum struct {
	COMPONENT CreateComponentRequestBodyKind
}

func GetCreateComponentRequestBodyKindEnum() CreateComponentRequestBodyKindEnum {
	return CreateComponentRequestBodyKindEnum{
		COMPONENT: CreateComponentRequestBodyKind{
			value: "Component",
		},
	}
}

func (c CreateComponentRequestBodyKind) Value() string {
	return c.value
}

func (c CreateComponentRequestBodyKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateComponentRequestBodyKind) UnmarshalJSON(b []byte) error {
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
