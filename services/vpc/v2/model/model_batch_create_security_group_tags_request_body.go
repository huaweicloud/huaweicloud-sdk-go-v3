package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateSecurityGroupTagsRequestBody This is a auto create Body Object
type BatchCreateSecurityGroupTagsRequestBody struct {

	// 操作标识
	Action BatchCreateSecurityGroupTagsRequestBodyAction `json:"action"`

	// 标签列表
	Tags []ResourceTag `json:"tags"`
}

func (o BatchCreateSecurityGroupTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSecurityGroupTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateSecurityGroupTagsRequestBody", string(data)}, " ")
}

type BatchCreateSecurityGroupTagsRequestBodyAction struct {
	value string
}

type BatchCreateSecurityGroupTagsRequestBodyActionEnum struct {
	CREATE BatchCreateSecurityGroupTagsRequestBodyAction
}

func GetBatchCreateSecurityGroupTagsRequestBodyActionEnum() BatchCreateSecurityGroupTagsRequestBodyActionEnum {
	return BatchCreateSecurityGroupTagsRequestBodyActionEnum{
		CREATE: BatchCreateSecurityGroupTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchCreateSecurityGroupTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchCreateSecurityGroupTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateSecurityGroupTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
