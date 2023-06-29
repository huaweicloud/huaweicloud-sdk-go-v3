package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateEnvironmentRequestBody struct {

	// API版本。
	ApiVersion CreateEnvironmentRequestBodyApiVersion `json:"api_version"`

	// 资源种类。
	Kind CreateEnvironmentRequestBodyKind `json:"kind"`

	Metadata *CreateEnvironmentRequestBodyMetadata `json:"metadata"`
}

func (o CreateEnvironmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentRequestBody", string(data)}, " ")
}

type CreateEnvironmentRequestBodyApiVersion struct {
	value string
}

type CreateEnvironmentRequestBodyApiVersionEnum struct {
	V1 CreateEnvironmentRequestBodyApiVersion
}

func GetCreateEnvironmentRequestBodyApiVersionEnum() CreateEnvironmentRequestBodyApiVersionEnum {
	return CreateEnvironmentRequestBodyApiVersionEnum{
		V1: CreateEnvironmentRequestBodyApiVersion{
			value: "v1",
		},
	}
}

func (c CreateEnvironmentRequestBodyApiVersion) Value() string {
	return c.value
}

func (c CreateEnvironmentRequestBodyApiVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEnvironmentRequestBodyApiVersion) UnmarshalJSON(b []byte) error {
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

type CreateEnvironmentRequestBodyKind struct {
	value string
}

type CreateEnvironmentRequestBodyKindEnum struct {
	ENVIRONMENT CreateEnvironmentRequestBodyKind
}

func GetCreateEnvironmentRequestBodyKindEnum() CreateEnvironmentRequestBodyKindEnum {
	return CreateEnvironmentRequestBodyKindEnum{
		ENVIRONMENT: CreateEnvironmentRequestBodyKind{
			value: "Environment",
		},
	}
}

func (c CreateEnvironmentRequestBodyKind) Value() string {
	return c.value
}

func (c CreateEnvironmentRequestBodyKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEnvironmentRequestBodyKind) UnmarshalJSON(b []byte) error {
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
