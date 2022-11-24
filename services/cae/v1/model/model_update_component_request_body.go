package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateComponentRequestBody struct {

	// API版本。
	ApiVersion UpdateComponentRequestBodyApiVersion `json:"api_version"`

	// 资源种类。
	Kind UpdateComponentRequestBodyKind `json:"kind"`

	Metadata *Metadata `json:"metadata,omitempty"`

	Spec *UpdateComponentRequestSpec `json:"spec,omitempty"`
}

func (o UpdateComponentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComponentRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateComponentRequestBody", string(data)}, " ")
}

type UpdateComponentRequestBodyApiVersion struct {
	value string
}

type UpdateComponentRequestBodyApiVersionEnum struct {
	V1 UpdateComponentRequestBodyApiVersion
}

func GetUpdateComponentRequestBodyApiVersionEnum() UpdateComponentRequestBodyApiVersionEnum {
	return UpdateComponentRequestBodyApiVersionEnum{
		V1: UpdateComponentRequestBodyApiVersion{
			value: "v1",
		},
	}
}

func (c UpdateComponentRequestBodyApiVersion) Value() string {
	return c.value
}

func (c UpdateComponentRequestBodyApiVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateComponentRequestBodyApiVersion) UnmarshalJSON(b []byte) error {
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

type UpdateComponentRequestBodyKind struct {
	value string
}

type UpdateComponentRequestBodyKindEnum struct {
	COMPONENT UpdateComponentRequestBodyKind
}

func GetUpdateComponentRequestBodyKindEnum() UpdateComponentRequestBodyKindEnum {
	return UpdateComponentRequestBodyKindEnum{
		COMPONENT: UpdateComponentRequestBodyKind{
			value: "Component",
		},
	}
}

func (c UpdateComponentRequestBodyKind) Value() string {
	return c.value
}

func (c UpdateComponentRequestBodyKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateComponentRequestBodyKind) UnmarshalJSON(b []byte) error {
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
