package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateComponentResponse struct {

	// API版本。
	ApiVersion *CreateComponentResponseApiVersion `json:"api_version,omitempty"`

	// 资源种类。
	Kind *CreateComponentResponseKind `json:"kind,omitempty"`

	Metadata *MetadataResponse `json:"metadata,omitempty"`

	Spec           *ComponentSpec `json:"spec,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateComponentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentResponse struct{}"
	}

	return strings.Join([]string{"CreateComponentResponse", string(data)}, " ")
}

type CreateComponentResponseApiVersion struct {
	value string
}

type CreateComponentResponseApiVersionEnum struct {
	V1 CreateComponentResponseApiVersion
}

func GetCreateComponentResponseApiVersionEnum() CreateComponentResponseApiVersionEnum {
	return CreateComponentResponseApiVersionEnum{
		V1: CreateComponentResponseApiVersion{
			value: "v1",
		},
	}
}

func (c CreateComponentResponseApiVersion) Value() string {
	return c.value
}

func (c CreateComponentResponseApiVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateComponentResponseApiVersion) UnmarshalJSON(b []byte) error {
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

type CreateComponentResponseKind struct {
	value string
}

type CreateComponentResponseKindEnum struct {
	COMPONENT CreateComponentResponseKind
}

func GetCreateComponentResponseKindEnum() CreateComponentResponseKindEnum {
	return CreateComponentResponseKindEnum{
		COMPONENT: CreateComponentResponseKind{
			value: "Component",
		},
	}
}

func (c CreateComponentResponseKind) Value() string {
	return c.value
}

func (c CreateComponentResponseKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateComponentResponseKind) UnmarshalJSON(b []byte) error {
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
