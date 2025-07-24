package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteTagsRequest Request Object
type BatchDeleteTagsRequest struct {

	// 资源类型
	ResourceType BatchDeleteTagsRequestResourceType `json:"resource_type"`

	// 资源id
	ResourceId string `json:"resource_id"`

	Body *ResourceTags `json:"body,omitempty"`
}

func (o BatchDeleteTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagsRequest", string(data)}, " ")
}

type BatchDeleteTagsRequestResourceType struct {
	value string
}

type BatchDeleteTagsRequestResourceTypeEnum struct {
	PHYSICALSERVERS BatchDeleteTagsRequestResourceType
}

func GetBatchDeleteTagsRequestResourceTypeEnum() BatchDeleteTagsRequestResourceTypeEnum {
	return BatchDeleteTagsRequestResourceTypeEnum{
		PHYSICALSERVERS: BatchDeleteTagsRequestResourceType{
			value: "physicalservers",
		},
	}
}

func (c BatchDeleteTagsRequestResourceType) Value() string {
	return c.value
}

func (c BatchDeleteTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
