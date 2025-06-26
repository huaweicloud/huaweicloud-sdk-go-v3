package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteSubResourceTagsRequest Request Object
type DeleteSubResourceTagsRequest struct {

	// 资源类型，支持的资源类型为：instances
	ResourceType DeleteSubResourceTagsRequestResourceType `json:"resource_type"`

	// 资源id
	ResourceId string `json:"resource_id"`

	// 子资源类型，支持的子资源类型为：namespaces
	SubResourceType DeleteSubResourceTagsRequestSubResourceType `json:"sub_resource_type"`

	// 子资源id
	SubResourceId string `json:"sub_resource_id"`

	Body *DeleteResourceTagsRequestBody `json:"body,omitempty"`
}

func (o DeleteSubResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubResourceTagsRequest", string(data)}, " ")
}

type DeleteSubResourceTagsRequestResourceType struct {
	value string
}

type DeleteSubResourceTagsRequestResourceTypeEnum struct {
	INSTANCES DeleteSubResourceTagsRequestResourceType
}

func GetDeleteSubResourceTagsRequestResourceTypeEnum() DeleteSubResourceTagsRequestResourceTypeEnum {
	return DeleteSubResourceTagsRequestResourceTypeEnum{
		INSTANCES: DeleteSubResourceTagsRequestResourceType{
			value: "instances",
		},
	}
}

func (c DeleteSubResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c DeleteSubResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteSubResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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

type DeleteSubResourceTagsRequestSubResourceType struct {
	value string
}

type DeleteSubResourceTagsRequestSubResourceTypeEnum struct {
	NAMESPACES DeleteSubResourceTagsRequestSubResourceType
}

func GetDeleteSubResourceTagsRequestSubResourceTypeEnum() DeleteSubResourceTagsRequestSubResourceTypeEnum {
	return DeleteSubResourceTagsRequestSubResourceTypeEnum{
		NAMESPACES: DeleteSubResourceTagsRequestSubResourceType{
			value: "namespaces",
		},
	}
}

func (c DeleteSubResourceTagsRequestSubResourceType) Value() string {
	return c.value
}

func (c DeleteSubResourceTagsRequestSubResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteSubResourceTagsRequestSubResourceType) UnmarshalJSON(b []byte) error {
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
