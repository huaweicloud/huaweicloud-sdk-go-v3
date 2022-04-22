package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 批量添加保护实例标签请求体
type BatchAddTagsRequestBody struct {

	// 标签列表。
	Tags []ResourceTag `json:"tags"`

	// 操作标识，取值仅限于：create：创建
	Action BatchAddTagsRequestBodyAction `json:"action"`
}

func (o BatchAddTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddTagsRequestBody", string(data)}, " ")
}

type BatchAddTagsRequestBodyAction struct {
	value string
}

type BatchAddTagsRequestBodyActionEnum struct {
	CREATE BatchAddTagsRequestBodyAction
}

func GetBatchAddTagsRequestBodyActionEnum() BatchAddTagsRequestBodyActionEnum {
	return BatchAddTagsRequestBodyActionEnum{
		CREATE: BatchAddTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchAddTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchAddTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
