package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// This is a auto create Body Object
type BatchCreateListenerTagsRequestBody struct {
	// 操作类型。 取值范围：create - 创建标签。

	Action BatchCreateListenerTagsRequestBodyAction `json:"action"`
	// 标签对象列表。

	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o BatchCreateListenerTagsRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateListenerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateListenerTagsRequestBody", string(data)}, " ")
}

type BatchCreateListenerTagsRequestBodyAction struct {
	value string
}

type BatchCreateListenerTagsRequestBodyActionEnum struct {
	CREATE BatchCreateListenerTagsRequestBodyAction
}

func GetBatchCreateListenerTagsRequestBodyActionEnum() BatchCreateListenerTagsRequestBodyActionEnum {
	return BatchCreateListenerTagsRequestBodyActionEnum{
		CREATE: BatchCreateListenerTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchCreateListenerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BatchCreateListenerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
