package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// This is a auto create Body Object
type BatchDeleteListenerTagsRequestBody struct {
	// 操作类型。 取值范围：delete- 删除标签。

	Action BatchDeleteListenerTagsRequestBodyAction `json:"action"`
	// 标签对象列表。

	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o BatchDeleteListenerTagsRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteListenerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteListenerTagsRequestBody", string(data)}, " ")
}

type BatchDeleteListenerTagsRequestBodyAction struct {
	value string
}

type BatchDeleteListenerTagsRequestBodyActionEnum struct {
	DELETE BatchDeleteListenerTagsRequestBodyAction
}

func GetBatchDeleteListenerTagsRequestBodyActionEnum() BatchDeleteListenerTagsRequestBodyActionEnum {
	return BatchDeleteListenerTagsRequestBodyActionEnum{
		DELETE: BatchDeleteListenerTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteListenerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BatchDeleteListenerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
