package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteListenerTagsRequestBody This is a auto create Body Object
type BatchDeleteListenerTagsRequestBody struct {

	// 操作类型。 取值范围：delete- 删除标签。
	Action BatchDeleteListenerTagsRequestBodyAction `json:"action"`

	// 标签对象列表。
	Tags []ResourceTag `json:"tags"`
}

func (o BatchDeleteListenerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
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

func (c BatchDeleteListenerTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchDeleteListenerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteListenerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
