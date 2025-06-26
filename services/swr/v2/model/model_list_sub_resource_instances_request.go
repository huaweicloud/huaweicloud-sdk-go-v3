package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSubResourceInstancesRequest Request Object
type ListSubResourceInstancesRequest struct {

	// 资源类型，支持的资源类型为：instances
	ResourceType ListSubResourceInstancesRequestResourceType `json:"resource_type"`

	// 资源id
	ResourceId string `json:"resource_id"`

	// 子资源类型，支持的子资源类型为：namespaces
	SubResourceType ListSubResourceInstancesRequestSubResourceType `json:"sub_resource_type"`

	// 起始索引，默认值为0。**注意：offset和limit参数需要配套使用。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为200，最小值为1，最大值为200。**注意：offset和limit参数需要配套使用。**
	Limit *int32 `json:"limit,omitempty"`

	Body *ListResourceInstancesRequestBody `json:"body,omitempty"`
}

func (o ListSubResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListSubResourceInstancesRequest", string(data)}, " ")
}

type ListSubResourceInstancesRequestResourceType struct {
	value string
}

type ListSubResourceInstancesRequestResourceTypeEnum struct {
	INSTANCES ListSubResourceInstancesRequestResourceType
}

func GetListSubResourceInstancesRequestResourceTypeEnum() ListSubResourceInstancesRequestResourceTypeEnum {
	return ListSubResourceInstancesRequestResourceTypeEnum{
		INSTANCES: ListSubResourceInstancesRequestResourceType{
			value: "instances",
		},
	}
}

func (c ListSubResourceInstancesRequestResourceType) Value() string {
	return c.value
}

func (c ListSubResourceInstancesRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSubResourceInstancesRequestResourceType) UnmarshalJSON(b []byte) error {
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

type ListSubResourceInstancesRequestSubResourceType struct {
	value string
}

type ListSubResourceInstancesRequestSubResourceTypeEnum struct {
	NAMESPACES ListSubResourceInstancesRequestSubResourceType
}

func GetListSubResourceInstancesRequestSubResourceTypeEnum() ListSubResourceInstancesRequestSubResourceTypeEnum {
	return ListSubResourceInstancesRequestSubResourceTypeEnum{
		NAMESPACES: ListSubResourceInstancesRequestSubResourceType{
			value: "namespaces",
		},
	}
}

func (c ListSubResourceInstancesRequestSubResourceType) Value() string {
	return c.value
}

func (c ListSubResourceInstancesRequestSubResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSubResourceInstancesRequestSubResourceType) UnmarshalJSON(b []byte) error {
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
