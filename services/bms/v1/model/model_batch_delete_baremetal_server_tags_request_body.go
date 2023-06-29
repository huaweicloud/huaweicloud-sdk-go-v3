package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteBaremetalServerTagsRequestBody This is a auto create Body Object
type BatchDeleteBaremetalServerTagsRequestBody struct {

	// 操作标识（仅支持小写）：delete（删除）。
	Action BatchDeleteBaremetalServerTagsRequestBodyAction `json:"action"`

	// 标签列表。
	Tags []BaremetalServerTag `json:"tags"`
}

func (o BatchDeleteBaremetalServerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
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

func (c BatchDeleteBaremetalServerTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchDeleteBaremetalServerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteBaremetalServerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
