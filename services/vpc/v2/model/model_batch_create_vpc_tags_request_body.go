package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateVpcTagsRequestBody This is a auto create Body Object
type BatchCreateVpcTagsRequestBody struct {

	// 功能说明：操作标识 取值范围：create
	Action BatchCreateVpcTagsRequestBodyAction `json:"action"`

	// 功能说明：标签列表
	Tags []ResourceTag `json:"tags"`
}

func (o BatchCreateVpcTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateVpcTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateVpcTagsRequestBody", string(data)}, " ")
}

type BatchCreateVpcTagsRequestBodyAction struct {
	value string
}

type BatchCreateVpcTagsRequestBodyActionEnum struct {
	CREATE BatchCreateVpcTagsRequestBodyAction
}

func GetBatchCreateVpcTagsRequestBodyActionEnum() BatchCreateVpcTagsRequestBodyActionEnum {
	return BatchCreateVpcTagsRequestBodyActionEnum{
		CREATE: BatchCreateVpcTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchCreateVpcTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchCreateVpcTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateVpcTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
