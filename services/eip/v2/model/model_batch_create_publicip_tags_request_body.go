package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreatePublicipTagsRequestBody 批量操作资源标签的请求体
type BatchCreatePublicipTagsRequestBody struct {

	// 标签列表
	Tags []ResourceTagOption `json:"tags"`

	// 操作标识  create：创建  action为create时，tag的value必选
	Action BatchCreatePublicipTagsRequestBodyAction `json:"action"`
}

func (o BatchCreatePublicipTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePublicipTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreatePublicipTagsRequestBody", string(data)}, " ")
}

type BatchCreatePublicipTagsRequestBodyAction struct {
	value string
}

type BatchCreatePublicipTagsRequestBodyActionEnum struct {
	CREATE BatchCreatePublicipTagsRequestBodyAction
}

func GetBatchCreatePublicipTagsRequestBodyActionEnum() BatchCreatePublicipTagsRequestBodyActionEnum {
	return BatchCreatePublicipTagsRequestBodyActionEnum{
		CREATE: BatchCreatePublicipTagsRequestBodyAction{
			value: "create",
		},
	}
}

func (c BatchCreatePublicipTagsRequestBodyAction) Value() string {
	return c.value
}

func (c BatchCreatePublicipTagsRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreatePublicipTagsRequestBodyAction) UnmarshalJSON(b []byte) error {
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
