package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
