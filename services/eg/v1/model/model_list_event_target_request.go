package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEventTargetRequest Request Object
type ListEventTargetRequest struct {

	// 偏移量，表示从此偏移量开始查询，偏移量不能小于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，不能小于1或大于1000
	Limit *int32 `json:"limit,omitempty"`

	// 指定查询排序
	Sort *string `json:"sort,omitempty"`

	// 指定查询的事件目标标签，模糊匹配
	FuzzyLabel *string `json:"fuzzy_label,omitempty"`

	// 事件目标支持方式：事件订阅：SUBSCRIPTION、事件流：FLOW
	SupportTypes *[]ListEventTargetRequestSupportTypes `json:"support_types,omitempty"`
}

func (o ListEventTargetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventTargetRequest struct{}"
	}

	return strings.Join([]string{"ListEventTargetRequest", string(data)}, " ")
}

type ListEventTargetRequestSupportTypes struct {
	value string
}

type ListEventTargetRequestSupportTypesEnum struct {
	SUBSCRIPTION ListEventTargetRequestSupportTypes
	FLOW         ListEventTargetRequestSupportTypes
}

func GetListEventTargetRequestSupportTypesEnum() ListEventTargetRequestSupportTypesEnum {
	return ListEventTargetRequestSupportTypesEnum{
		SUBSCRIPTION: ListEventTargetRequestSupportTypes{
			value: "SUBSCRIPTION",
		},
		FLOW: ListEventTargetRequestSupportTypes{
			value: "FLOW",
		},
	}
}

func (c ListEventTargetRequestSupportTypes) Value() string {
	return c.value
}

func (c ListEventTargetRequestSupportTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventTargetRequestSupportTypes) UnmarshalJSON(b []byte) error {
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
