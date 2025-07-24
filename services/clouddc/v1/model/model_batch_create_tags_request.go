package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateTagsRequest Request Object
type BatchCreateTagsRequest struct {

	// 资源类型
	ResourceType BatchCreateTagsRequestResourceType `json:"resource_type"`

	// 资源id
	ResourceId string `json:"resource_id"`

	Body *ResourceTags `json:"body,omitempty"`
}

func (o BatchCreateTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateTagsRequest", string(data)}, " ")
}

type BatchCreateTagsRequestResourceType struct {
	value string
}

type BatchCreateTagsRequestResourceTypeEnum struct {
	PHYSICALSERVERS BatchCreateTagsRequestResourceType
}

func GetBatchCreateTagsRequestResourceTypeEnum() BatchCreateTagsRequestResourceTypeEnum {
	return BatchCreateTagsRequestResourceTypeEnum{
		PHYSICALSERVERS: BatchCreateTagsRequestResourceType{
			value: "physicalservers",
		},
	}
}

func (c BatchCreateTagsRequestResourceType) Value() string {
	return c.value
}

func (c BatchCreateTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
