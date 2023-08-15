package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateServerTagsRequestBody This is a auto create Body Object
type BatchCreateServerTagsRequestBody struct {

	// 操作标识（仅支持小写）：create（创建）。
	Action BatchCreateServerTagsRequestBodyAction `json:"action"`

	// 标签列表。
	Tags []ServerTag `json:"tags"`
}

func (o BatchCreateServerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateServerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateServerTagsRequestBody", string(data)}, " ")
}

type BatchCreateServerTagsRequestBodyAction struct {
	value string
}

type BatchCreateServerTagsRequestBodyActionEnum struct {
	CREATE BatchCreateServerTagsRequestBodyAction
}

func GetBatchCreateServerTagsRequestBodyActionEnum() BatchCreateServerTagsRequestBodyActionEnum {
	return BatchCreateServerTagsRequestBodyActionEnum{
		CREATE: BatchCreateServerTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchCreateServerTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchCreateServerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateServerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
