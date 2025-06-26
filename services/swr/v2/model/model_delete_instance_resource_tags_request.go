package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteInstanceResourceTagsRequest Request Object
type DeleteInstanceResourceTagsRequest struct {

	// 资源类型，支持的资源类型为：instances
	ResourceType DeleteInstanceResourceTagsRequestResourceType `json:"resource_type"`

	// 资源id
	ResourceId string `json:"resource_id"`

	Body *DeleteResourceTagsRequestBody `json:"body,omitempty"`
}

func (o DeleteInstanceResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceResourceTagsRequest", string(data)}, " ")
}

type DeleteInstanceResourceTagsRequestResourceType struct {
	value string
}

type DeleteInstanceResourceTagsRequestResourceTypeEnum struct {
	INSTANCES DeleteInstanceResourceTagsRequestResourceType
}

func GetDeleteInstanceResourceTagsRequestResourceTypeEnum() DeleteInstanceResourceTagsRequestResourceTypeEnum {
	return DeleteInstanceResourceTagsRequestResourceTypeEnum{
		INSTANCES: DeleteInstanceResourceTagsRequestResourceType{
			value: "instances",
		},
	}
}

func (c DeleteInstanceResourceTagsRequestResourceType) Value() string {
	return c.value
}

func (c DeleteInstanceResourceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteInstanceResourceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
