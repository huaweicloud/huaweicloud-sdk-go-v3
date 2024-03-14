package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchAddOrRemoveResourceInstanceRequestBody 批量添加或删除资源标签接口请求结构体
type BatchAddOrRemoveResourceInstanceRequestBody struct {

	// 标签列表，没有标签默认为空数组。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 操作标识：仅限于 create（创建） delete（删除）
	Action BatchAddOrRemoveResourceInstanceRequestBodyAction `json:"action"`
}

func (o BatchAddOrRemoveResourceInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddOrRemoveResourceInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddOrRemoveResourceInstanceRequestBody", string(data)}, " ")
}

type BatchAddOrRemoveResourceInstanceRequestBodyAction struct {
	value string
}

type BatchAddOrRemoveResourceInstanceRequestBodyActionEnum struct {
	CREATE BatchAddOrRemoveResourceInstanceRequestBodyAction
	DELETE BatchAddOrRemoveResourceInstanceRequestBodyAction
}

func GetBatchAddOrRemoveResourceInstanceRequestBodyActionEnum() BatchAddOrRemoveResourceInstanceRequestBodyActionEnum {
	return BatchAddOrRemoveResourceInstanceRequestBodyActionEnum{
		CREATE: BatchAddOrRemoveResourceInstanceRequestBodyAction{
			value: "create",
		},
		DELETE: BatchAddOrRemoveResourceInstanceRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchAddOrRemoveResourceInstanceRequestBodyAction) Value() string {
	return c.value
}

func (c BatchAddOrRemoveResourceInstanceRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchAddOrRemoveResourceInstanceRequestBodyAction) UnmarshalJSON(b []byte) error {
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
