package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteLoadbalancerTagsRequestBody This is a auto create Body Object
type BatchDeleteLoadbalancerTagsRequestBody struct {

	// 操作类型。 取值范围：delete - 删除标签。
	Action BatchDeleteLoadbalancerTagsRequestBodyAction `json:"action"`

	// 标签对象列表。
	Tags []ResourceTag `json:"tags"`
}

func (o BatchDeleteLoadbalancerTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteLoadbalancerTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteLoadbalancerTagsRequestBody", string(data)}, " ")
}

type BatchDeleteLoadbalancerTagsRequestBodyAction struct {
	value string
}

type BatchDeleteLoadbalancerTagsRequestBodyActionEnum struct {
	DELETE BatchDeleteLoadbalancerTagsRequestBodyAction
}

func GetBatchDeleteLoadbalancerTagsRequestBodyActionEnum() BatchDeleteLoadbalancerTagsRequestBodyActionEnum {
	return BatchDeleteLoadbalancerTagsRequestBodyActionEnum{
		DELETE: BatchDeleteLoadbalancerTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteLoadbalancerTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchDeleteLoadbalancerTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteLoadbalancerTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
