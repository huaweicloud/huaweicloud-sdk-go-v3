package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSubResourceTagsRequest Request Object
type ListSubResourceTagsRequest struct {

	// 资源类型，支持的资源类型为：instances
	ResourceType ListSubResourceTagsRequestResourceType `json:"resource_type"`

	// 资源id
	ResourceId string `json:"resource_id"`

	// 子资源类型，支持的子资源类型为：namespaces
	SubResourceType ListSubResourceTagsRequestSubResourceType `json:"sub_resource_type"`

	// 子资源id
	SubResourceId string `json:"sub_resource_id"`
}

func (o ListSubResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSubResourceTagsRequest", string(data)}, " ")
}

type ListSubResourceTagsRequestResourceType struct {
	value string
}

type ListSubResourceTagsRequestResourceTypeEnum struct {
	INSTANCES ListSubResourceTagsRequestResourceType
}

func GetListSubResourceTagsRequestResourceTypeEnum() ListSubResourceTagsRequestResourceTypeEnum {
	return ListSubResourceTagsRequestResourceTypeEnum{
		INSTANCES: ListSubResourceTagsRequestResourceType{
			value: "instances",
		},
	}
}

func (c ListSubResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListSubResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSubResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type ListSubResourceTagsRequestSubResourceType struct {
	value string
}

type ListSubResourceTagsRequestSubResourceTypeEnum struct {
	NAMESPACES ListSubResourceTagsRequestSubResourceType
}

func GetListSubResourceTagsRequestSubResourceTypeEnum() ListSubResourceTagsRequestSubResourceTypeEnum {
	return ListSubResourceTagsRequestSubResourceTypeEnum{
		NAMESPACES: ListSubResourceTagsRequestSubResourceType{
			value: "namespaces",
		},
	}
}

func (c ListSubResourceTagsRequestSubResourceType) Value() string {
	return c.value
}

func (c ListSubResourceTagsRequestSubResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSubResourceTagsRequestSubResourceType) UnmarshalJSON(b []byte) error {
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
