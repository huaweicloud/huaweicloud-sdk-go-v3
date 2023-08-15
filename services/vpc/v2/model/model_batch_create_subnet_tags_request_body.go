package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateSubnetTagsRequestBody This is a auto create Body Object
type BatchCreateSubnetTagsRequestBody struct {

	// 操作标识
	Action BatchCreateSubnetTagsRequestBodyAction `json:"action"`

	// 标签列表
	Tags []ResourceTag `json:"tags"`
}

func (o BatchCreateSubnetTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSubnetTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateSubnetTagsRequestBody", string(data)}, " ")
}

type BatchCreateSubnetTagsRequestBodyAction struct {
	value string
}

type BatchCreateSubnetTagsRequestBodyActionEnum struct {
	CREATE BatchCreateSubnetTagsRequestBodyAction
}

func GetBatchCreateSubnetTagsRequestBodyActionEnum() BatchCreateSubnetTagsRequestBodyActionEnum {
	return BatchCreateSubnetTagsRequestBodyActionEnum{
		CREATE: BatchCreateSubnetTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchCreateSubnetTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchCreateSubnetTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateSubnetTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
