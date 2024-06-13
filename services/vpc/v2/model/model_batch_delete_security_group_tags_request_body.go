package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteSecurityGroupTagsRequestBody This is a auto create Body Object
type BatchDeleteSecurityGroupTagsRequestBody struct {

	// 功能说明：操作标识 取值范围：delete
	Action BatchDeleteSecurityGroupTagsRequestBodyAction `json:"action"`

	// 标签列表
	Tags []ResourceTag `json:"tags"`
}

func (o BatchDeleteSecurityGroupTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecurityGroupTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecurityGroupTagsRequestBody", string(data)}, " ")
}

type BatchDeleteSecurityGroupTagsRequestBodyAction struct {
	value string
}

type BatchDeleteSecurityGroupTagsRequestBodyActionEnum struct {
	DELETE BatchDeleteSecurityGroupTagsRequestBodyAction
}

func GetBatchDeleteSecurityGroupTagsRequestBodyActionEnum() BatchDeleteSecurityGroupTagsRequestBodyActionEnum {
	return BatchDeleteSecurityGroupTagsRequestBodyActionEnum{
		DELETE: BatchDeleteSecurityGroupTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteSecurityGroupTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchDeleteSecurityGroupTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteSecurityGroupTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
