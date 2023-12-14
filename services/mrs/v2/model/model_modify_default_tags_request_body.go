package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyDefaultTagsRequestBody 操作默认标签请求体
type ModifyDefaultTagsRequestBody struct {

	// 操作类型，支持创建和删除
	Action ModifyDefaultTagsRequestBodyAction `json:"action"`
}

func (o ModifyDefaultTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDefaultTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyDefaultTagsRequestBody", string(data)}, " ")
}

type ModifyDefaultTagsRequestBodyAction struct {
	value string
}

type ModifyDefaultTagsRequestBodyActionEnum struct {
	CREATE ModifyDefaultTagsRequestBodyAction
	DELETE ModifyDefaultTagsRequestBodyAction
}

func GetModifyDefaultTagsRequestBodyActionEnum() ModifyDefaultTagsRequestBodyActionEnum {
	return ModifyDefaultTagsRequestBodyActionEnum{
		CREATE: ModifyDefaultTagsRequestBodyAction{
			value: "create",
		},
		DELETE: ModifyDefaultTagsRequestBodyAction{
			value: "delete",
		},
	}
}

func (c ModifyDefaultTagsRequestBodyAction) Value() string {
	return c.value
}

func (c ModifyDefaultTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyDefaultTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
