package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateListenerTagsRequestBody This is a auto create Body Object
type BatchCreateListenerTagsRequestBody struct {

	// 操作类型。 取值范围：create - 创建标签。
	Action BatchCreateListenerTagsRequestBodyAction `json:"action"`

	// 标签对象列表。
	Tags []ResourceTag `json:"tags"`
}

func (o BatchCreateListenerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
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

func (c BatchCreateListenerTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchCreateListenerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateListenerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
