package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateTagsRequest Request Object
type BatchCreateTagsRequest struct {

	// 资源类型。  - cph-server，云手机服务器
	ResourceType BatchCreateTagsRequestResourceType `json:"resource_type"`

	// 资源ID。
	ResourceId string `json:"resource_id"`

	Body *BatchCreateTagsRequestBody `json:"body,omitempty"`
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
	CPH_SERVER BatchCreateTagsRequestResourceType
}

func GetBatchCreateTagsRequestResourceTypeEnum() BatchCreateTagsRequestResourceTypeEnum {
	return BatchCreateTagsRequestResourceTypeEnum{
		CPH_SERVER: BatchCreateTagsRequestResourceType{
			value: "cph-server",
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
