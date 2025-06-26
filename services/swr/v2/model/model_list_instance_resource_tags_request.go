package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceResourceTagsRequest Request Object
type ListInstanceResourceTagsRequest struct {

	// 资源类型，支持的资源类型为：instances
	ResourceType ListInstanceResourceTagsRequestResourceType `json:"resource_type"`

	// 资源id
	ResourceId string `json:"resource_id"`
}

func (o ListInstanceResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceResourceTagsRequest", string(data)}, " ")
}

type ListInstanceResourceTagsRequestResourceType struct {
	value string
}

type ListInstanceResourceTagsRequestResourceTypeEnum struct {
	INSTANCES ListInstanceResourceTagsRequestResourceType
}

func GetListInstanceResourceTagsRequestResourceTypeEnum() ListInstanceResourceTagsRequestResourceTypeEnum {
	return ListInstanceResourceTagsRequestResourceTypeEnum{
		INSTANCES: ListInstanceResourceTagsRequestResourceType{
			value: "instances",
		},
	}
}

func (c ListInstanceResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListInstanceResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
