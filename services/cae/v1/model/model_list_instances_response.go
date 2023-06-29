package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstancesResponse Response Object
type ListInstancesResponse struct {

	// API版本。
	ApiVersion *ListInstancesResponseApiVersion `json:"api_version,omitempty"`

	// 资源种类。
	Kind *ListInstancesResponseKind `json:"kind,omitempty"`

	// 组件实例列表。
	Items          *[]Instance `json:"items,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesResponse", string(data)}, " ")
}

type ListInstancesResponseApiVersion struct {
	value string
}

type ListInstancesResponseApiVersionEnum struct {
	V1 ListInstancesResponseApiVersion
}

func GetListInstancesResponseApiVersionEnum() ListInstancesResponseApiVersionEnum {
	return ListInstancesResponseApiVersionEnum{
		V1: ListInstancesResponseApiVersion{
			value: "v1",
		},
	}
}

func (c ListInstancesResponseApiVersion) Value() string {
	return c.value
}

func (c ListInstancesResponseApiVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesResponseApiVersion) UnmarshalJSON(b []byte) error {
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

type ListInstancesResponseKind struct {
	value string
}

type ListInstancesResponseKindEnum struct {
	COMPONENT_INSTANCE ListInstancesResponseKind
}

func GetListInstancesResponseKindEnum() ListInstancesResponseKindEnum {
	return ListInstancesResponseKindEnum{
		COMPONENT_INSTANCE: ListInstancesResponseKind{
			value: "ComponentInstance",
		},
	}
}

func (c ListInstancesResponseKind) Value() string {
	return c.value
}

func (c ListInstancesResponseKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesResponseKind) UnmarshalJSON(b []byte) error {
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
