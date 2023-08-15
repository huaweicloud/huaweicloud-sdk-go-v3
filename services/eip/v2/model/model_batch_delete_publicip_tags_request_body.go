package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeletePublicipTagsRequestBody 批量操作资源标签的请求体
type BatchDeletePublicipTagsRequestBody struct {

	// 标签列表
	Tags []ResourceTagOption `json:"tags"`

	// 操作标识  delete：删除  action为delete时，value可选
	Action BatchDeletePublicipTagsRequestBodyAction `json:"action"`
}

func (o BatchDeletePublicipTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePublicipTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeletePublicipTagsRequestBody", string(data)}, " ")
}

type BatchDeletePublicipTagsRequestBodyAction struct {
	value string
}

type BatchDeletePublicipTagsRequestBodyActionEnum struct {
	DELETE BatchDeletePublicipTagsRequestBodyAction
}

func GetBatchDeletePublicipTagsRequestBodyActionEnum() BatchDeletePublicipTagsRequestBodyActionEnum {
	return BatchDeletePublicipTagsRequestBodyActionEnum{
		DELETE: BatchDeletePublicipTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeletePublicipTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchDeletePublicipTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeletePublicipTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
