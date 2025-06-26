package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSubResourceTagsRequest Request Object
type CreateSubResourceTagsRequest struct {

	// 资源类型，支持的资源类型为：instances
	ResourceType CreateSubResourceTagsRequestResourceType `json:"resource_type"`

	// 资源id
	ResourceId string `json:"resource_id"`

	// 子资源类型，支持的子资源类型为：namespaces
	SubResourceType CreateSubResourceTagsRequestSubResourceType `json:"sub_resource_type"`

	// 子资源id
	SubResourceId string `json:"sub_resource_id"`

	Body *CreateResourceTagsRequestBody `json:"body,omitempty"`
}

func (o CreateSubResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateSubResourceTagsRequest", string(data)}, " ")
}

type CreateSubResourceTagsRequestResourceType struct {
	value string
}

type CreateSubResourceTagsRequestResourceTypeEnum struct {
	INSTANCES CreateSubResourceTagsRequestResourceType
}

func GetCreateSubResourceTagsRequestResourceTypeEnum() CreateSubResourceTagsRequestResourceTypeEnum {
	return CreateSubResourceTagsRequestResourceTypeEnum{
		INSTANCES: CreateSubResourceTagsRequestResourceType{
			value: "instances",
		},
	}
}

func (c CreateSubResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c CreateSubResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSubResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type CreateSubResourceTagsRequestSubResourceType struct {
	value string
}

type CreateSubResourceTagsRequestSubResourceTypeEnum struct {
	NAMESPACES CreateSubResourceTagsRequestSubResourceType
}

func GetCreateSubResourceTagsRequestSubResourceTypeEnum() CreateSubResourceTagsRequestSubResourceTypeEnum {
	return CreateSubResourceTagsRequestSubResourceTypeEnum{
		NAMESPACES: CreateSubResourceTagsRequestSubResourceType{
			value: "namespaces",
		},
	}
}

func (c CreateSubResourceTagsRequestSubResourceType) Value() string {
	return c.value
}

func (c CreateSubResourceTagsRequestSubResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSubResourceTagsRequestSubResourceType) UnmarshalJSON(b []byte) error {
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
