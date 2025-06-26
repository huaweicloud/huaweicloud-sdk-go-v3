package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateInstanceResourceTagsRequest Request Object
type CreateInstanceResourceTagsRequest struct {

	// 资源类型，支持的资源类型为：instances
	ResourceType CreateInstanceResourceTagsRequestResourceType `json:"resource_type"`

	// 资源id
	ResourceId string `json:"resource_id"`

	Body *CreateResourceTagsRequestBody `json:"body,omitempty"`
}

func (o CreateInstanceResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceResourceTagsRequest", string(data)}, " ")
}

type CreateInstanceResourceTagsRequestResourceType struct {
	value string
}

type CreateInstanceResourceTagsRequestResourceTypeEnum struct {
	INSTANCES CreateInstanceResourceTagsRequestResourceType
}

func GetCreateInstanceResourceTagsRequestResourceTypeEnum() CreateInstanceResourceTagsRequestResourceTypeEnum {
	return CreateInstanceResourceTagsRequestResourceTypeEnum{
		INSTANCES: CreateInstanceResourceTagsRequestResourceType{
			value: "instances",
		},
	}
}

func (c CreateInstanceResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c CreateInstanceResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
