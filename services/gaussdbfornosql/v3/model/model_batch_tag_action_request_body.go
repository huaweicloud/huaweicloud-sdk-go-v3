package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type BatchTagActionRequestBody struct {
	// 操作标识。取值： - create，表示添加标签。 - delete，表示删除标签。

	Action BatchTagActionRequestBodyAction `json:"action"`
	// 标签列表。

	Tags []BatchTagActionTagOption `json:"tags"`
}

func (o BatchTagActionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagActionRequestBody struct{}"
	}

	return strings.Join([]string{"BatchTagActionRequestBody", string(data)}, " ")
}

type BatchTagActionRequestBodyAction struct {
	value string
}

type BatchTagActionRequestBodyActionEnum struct {
	CREATE BatchTagActionRequestBodyAction
	DELETE BatchTagActionRequestBodyAction
}

func GetBatchTagActionRequestBodyActionEnum() BatchTagActionRequestBodyActionEnum {
	return BatchTagActionRequestBodyActionEnum{
		CREATE: BatchTagActionRequestBodyAction{
			value: "create",
		},
		DELETE: BatchTagActionRequestBodyAction{
			value: "delete",
		},
	}
}

func (c BatchTagActionRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchTagActionRequestBodyAction) UnmarshalJSON(b []byte) error {
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
