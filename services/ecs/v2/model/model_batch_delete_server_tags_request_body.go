package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteServerTagsRequestBody This is a auto create Body Object
type BatchDeleteServerTagsRequestBody struct {

	// 操作标识（仅支持小写）：delete（删除）。
	Action BatchDeleteServerTagsRequestBodyAction `json:"action"`

	// 标签列表。
	Tags []ServerTag `json:"tags"`
}

func (o BatchDeleteServerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerTagsRequestBody", string(data)}, " ")
}

type BatchDeleteServerTagsRequestBodyAction struct {
	value string
}

type BatchDeleteServerTagsRequestBodyActionEnum struct {
	DELETE BatchDeleteServerTagsRequestBodyAction
}

func GetBatchDeleteServerTagsRequestBodyActionEnum() BatchDeleteServerTagsRequestBodyActionEnum {
	return BatchDeleteServerTagsRequestBodyActionEnum{
		DELETE: BatchDeleteServerTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteServerTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchDeleteServerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteServerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
