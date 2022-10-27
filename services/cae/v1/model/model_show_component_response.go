package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowComponentResponse struct {

	// API版本。
	ApiVersion *ShowComponentResponseApiVersion `json:"api_version,omitempty"`

	// 资源种类。
	Kind *ShowComponentResponseKind `json:"kind,omitempty"`

	Metadata *MetadataResponse `json:"metadata,omitempty"`

	Spec           *ComponentSpec `json:"spec,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowComponentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComponentResponse struct{}"
	}

	return strings.Join([]string{"ShowComponentResponse", string(data)}, " ")
}

type ShowComponentResponseApiVersion struct {
	value string
}

type ShowComponentResponseApiVersionEnum struct {
	V1 ShowComponentResponseApiVersion
}

func GetShowComponentResponseApiVersionEnum() ShowComponentResponseApiVersionEnum {
	return ShowComponentResponseApiVersionEnum{
		V1: ShowComponentResponseApiVersion{
			value: "v1",
		},
	}
}

func (c ShowComponentResponseApiVersion) Value() string {
	return c.value
}

func (c ShowComponentResponseApiVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowComponentResponseApiVersion) UnmarshalJSON(b []byte) error {
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

type ShowComponentResponseKind struct {
	value string
}

type ShowComponentResponseKindEnum struct {
	COMPONENT ShowComponentResponseKind
}

func GetShowComponentResponseKindEnum() ShowComponentResponseKindEnum {
	return ShowComponentResponseKindEnum{
		COMPONENT: ShowComponentResponseKind{
			value: "Component",
		},
	}
}

func (c ShowComponentResponseKind) Value() string {
	return c.value
}

func (c ShowComponentResponseKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowComponentResponseKind) UnmarshalJSON(b []byte) error {
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
