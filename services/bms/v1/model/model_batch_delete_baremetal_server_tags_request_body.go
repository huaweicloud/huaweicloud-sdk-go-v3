package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// This is a auto create Body Object
type BatchDeleteBaremetalServerTagsRequestBody struct {
	// 操作标识（仅支持小写）：delete（删除）。

	Action BatchDeleteBaremetalServerTagsRequestBodyAction `json:"action"`
	// 标签列表。

	Tags []BaremetalServerTag `json:"tags"`
}

func (o BatchDeleteBaremetalServerTagsRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteBaremetalServerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteBaremetalServerTagsRequestBody", string(data)}, " ")
}

type BatchDeleteBaremetalServerTagsRequestBodyAction struct {
	value string
}

type BatchDeleteBaremetalServerTagsRequestBodyActionEnum struct {
	DELETE BatchDeleteBaremetalServerTagsRequestBodyAction
}

func GetBatchDeleteBaremetalServerTagsRequestBodyActionEnum() BatchDeleteBaremetalServerTagsRequestBodyActionEnum {
	return BatchDeleteBaremetalServerTagsRequestBodyActionEnum{
		DELETE: BatchDeleteBaremetalServerTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteBaremetalServerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BatchDeleteBaremetalServerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
