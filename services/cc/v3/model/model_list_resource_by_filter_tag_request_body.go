package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResourceByFilterTagRequestBody 用于查询资源实例的Tag过滤条件
type ListResourceByFilterTagRequestBody struct {

	// 动作。|- filter：过滤。 count：查询总条数。
	Action *ListResourceByFilterTagRequestBodyAction `json:"action,omitempty"`

	// 查询结果数量限制
	Limit *int32 `json:"limit,omitempty"`

	// 查询结果偏移
	Offset *int32 `json:"offset,omitempty"`

	// 是否包含以下tag（多个key取\"与\"关系，多个value取\"或\"关系）
	Tags *[]AggTag `json:"tags,omitempty"`

	// 是否匹配以下tag，key必须为\"resource_name\"，value如果有值则模糊匹配，如果为空字符串则精确匹配
	Matches *[]Tag `json:"matches,omitempty"`
}

func (o ListResourceByFilterTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceByFilterTagRequestBody struct{}"
	}

	return strings.Join([]string{"ListResourceByFilterTagRequestBody", string(data)}, " ")
}

type ListResourceByFilterTagRequestBodyAction struct {
	value string
}

type ListResourceByFilterTagRequestBodyActionEnum struct {
	FILTER ListResourceByFilterTagRequestBodyAction
	COUNT  ListResourceByFilterTagRequestBodyAction
}

func GetListResourceByFilterTagRequestBodyActionEnum() ListResourceByFilterTagRequestBodyActionEnum {
	return ListResourceByFilterTagRequestBodyActionEnum{
		FILTER: ListResourceByFilterTagRequestBodyAction{
			value: "filter",
		},
		COUNT: ListResourceByFilterTagRequestBodyAction{
			value: "count",
		},
	}
}

func (c ListResourceByFilterTagRequestBodyAction) Value() string {
	return c.value
}

func (c ListResourceByFilterTagRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResourceByFilterTagRequestBodyAction) UnmarshalJSON(b []byte) error {
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
