package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateBaremetalServerTagsRequestBody This is a auto create Body Object
type BatchCreateBaremetalServerTagsRequestBody struct {

	// 操作标识（仅支持小写）：create（创建）。
	Action BatchCreateBaremetalServerTagsRequestBodyAction `json:"action"`

	// 标签列表。
	Tags []BaremetalServerTag `json:"tags"`
}

func (o BatchCreateBaremetalServerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateBaremetalServerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateBaremetalServerTagsRequestBody", string(data)}, " ")
}

type BatchCreateBaremetalServerTagsRequestBodyAction struct {
	value string
}

type BatchCreateBaremetalServerTagsRequestBodyActionEnum struct {
	CREATE BatchCreateBaremetalServerTagsRequestBodyAction
}

func GetBatchCreateBaremetalServerTagsRequestBodyActionEnum() BatchCreateBaremetalServerTagsRequestBodyActionEnum {
	return BatchCreateBaremetalServerTagsRequestBodyActionEnum{
		CREATE: BatchCreateBaremetalServerTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchCreateBaremetalServerTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchCreateBaremetalServerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateBaremetalServerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
