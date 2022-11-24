package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateApplicationRequestBody struct {

	// API版本。
	ApiVersion CreateApplicationRequestBodyApiVersion `json:"api_version"`

	// 资源种类。
	Kind CreateApplicationRequestBodyKind `json:"kind"`

	Metadata *CreateApplicationRequestBodyMetadata `json:"metadata"`
}

func (o CreateApplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateApplicationRequestBody", string(data)}, " ")
}

type CreateApplicationRequestBodyApiVersion struct {
	value string
}

type CreateApplicationRequestBodyApiVersionEnum struct {
	V1 CreateApplicationRequestBodyApiVersion
}

func GetCreateApplicationRequestBodyApiVersionEnum() CreateApplicationRequestBodyApiVersionEnum {
	return CreateApplicationRequestBodyApiVersionEnum{
		V1: CreateApplicationRequestBodyApiVersion{
			value: "v1",
		},
	}
}

func (c CreateApplicationRequestBodyApiVersion) Value() string {
	return c.value
}

func (c CreateApplicationRequestBodyApiVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateApplicationRequestBodyApiVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateApplicationRequestBodyKind struct {
	value string
}

type CreateApplicationRequestBodyKindEnum struct {
	APPLICATION CreateApplicationRequestBodyKind
}

func GetCreateApplicationRequestBodyKindEnum() CreateApplicationRequestBodyKindEnum {
	return CreateApplicationRequestBodyKindEnum{
		APPLICATION: CreateApplicationRequestBodyKind{
			value: "Application",
		},
	}
}

func (c CreateApplicationRequestBodyKind) Value() string {
	return c.value
}

func (c CreateApplicationRequestBodyKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateApplicationRequestBodyKind) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
