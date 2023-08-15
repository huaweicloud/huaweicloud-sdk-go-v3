package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteTagsRequestBody 批量删除保护实例标签请求体
type BatchDeleteTagsRequestBody struct {

	// 标签列表。
	Tags []DeleteResourceTag `json:"tags"`

	// 操作标识，取值仅限于：delete：删除
	Action BatchDeleteTagsRequestBodyAction `json:"action"`
}

func (o BatchDeleteTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteTagsRequestBody", string(data)}, " ")
}

type BatchDeleteTagsRequestBodyAction struct {
	value string
}

type BatchDeleteTagsRequestBodyActionEnum struct {
	DELETE BatchDeleteTagsRequestBodyAction
}

func GetBatchDeleteTagsRequestBodyActionEnum() BatchDeleteTagsRequestBodyActionEnum {
	return BatchDeleteTagsRequestBodyActionEnum{
		DELETE: BatchDeleteTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchDeleteTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
