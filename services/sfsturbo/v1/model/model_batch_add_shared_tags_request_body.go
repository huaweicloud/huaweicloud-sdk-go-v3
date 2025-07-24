package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchAddSharedTagsRequestBody struct {

	// 操作标识，取值范围为：create。 为指定共享批量添加标签时使用create。
	Action BatchAddSharedTagsRequestBodyAction `json:"action"`

	// 标签列表。
	Tags []ResourceTag `json:"tags"`
}

func (o BatchAddSharedTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddSharedTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddSharedTagsRequestBody", string(data)}, " ")
}

type BatchAddSharedTagsRequestBodyAction struct {
	value string
}

type BatchAddSharedTagsRequestBodyActionEnum struct {
	CREATE BatchAddSharedTagsRequestBodyAction
}

func GetBatchAddSharedTagsRequestBodyActionEnum() BatchAddSharedTagsRequestBodyActionEnum {
	return BatchAddSharedTagsRequestBodyActionEnum{
		CREATE: BatchAddSharedTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchAddSharedTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchAddSharedTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchAddSharedTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
