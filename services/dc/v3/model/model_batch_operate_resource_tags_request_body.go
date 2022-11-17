package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 批量添加删除的请求体。
type BatchOperateResourceTagsRequestBody struct {

	// 功能说明：操作标识。 取值范围： create（创建） delete（删除）
	Action BatchOperateResourceTagsRequestBodyAction `json:"action"`

	// 标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// 标签列表。
	SysTags *[]Tag `json:"sys_tags,omitempty"`
}

func (o BatchOperateResourceTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateResourceTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchOperateResourceTagsRequestBody", string(data)}, " ")
}

type BatchOperateResourceTagsRequestBodyAction struct {
	value string
}

type BatchOperateResourceTagsRequestBodyActionEnum struct {
	CREATE BatchOperateResourceTagsRequestBodyAction
	DELETE BatchOperateResourceTagsRequestBodyAction
}

func GetBatchOperateResourceTagsRequestBodyActionEnum() BatchOperateResourceTagsRequestBodyActionEnum {
	return BatchOperateResourceTagsRequestBodyActionEnum{
		CREATE: BatchOperateResourceTagsRequestBodyAction{
			value: "create",
		},
		DELETE: BatchOperateResourceTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchOperateResourceTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchOperateResourceTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchOperateResourceTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
