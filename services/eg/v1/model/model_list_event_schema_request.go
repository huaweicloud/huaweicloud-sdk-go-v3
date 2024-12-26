package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEventSchemaRequest Request Object
type ListEventSchemaRequest struct {

	// 偏移量，表示从此偏移量开始查询，偏移量不能小于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，不能小于1或大于1000
	Limit *int32 `json:"limit,omitempty"`

	// 指定查询排序
	Sort *string `json:"sort,omitempty"`

	// 指定查询提供方的类型
	ProviderType *ListEventSchemaRequestProviderType `json:"provider_type,omitempty"`

	// 指定查询的事件模型名称，精准匹配
	Name *string `json:"name,omitempty"`

	// 指定查询的事件模型名称，模糊匹配
	FuzzyName *string `json:"fuzzy_name,omitempty"`
}

func (o ListEventSchemaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventSchemaRequest struct{}"
	}

	return strings.Join([]string{"ListEventSchemaRequest", string(data)}, " ")
}

type ListEventSchemaRequestProviderType struct {
	value string
}

type ListEventSchemaRequestProviderTypeEnum struct {
	OFFICIAL ListEventSchemaRequestProviderType
	CUSTOM   ListEventSchemaRequestProviderType
	PARTNER  ListEventSchemaRequestProviderType
}

func GetListEventSchemaRequestProviderTypeEnum() ListEventSchemaRequestProviderTypeEnum {
	return ListEventSchemaRequestProviderTypeEnum{
		OFFICIAL: ListEventSchemaRequestProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: ListEventSchemaRequestProviderType{
			value: "CUSTOM",
		},
		PARTNER: ListEventSchemaRequestProviderType{
			value: "PARTNER",
		},
	}
}

func (c ListEventSchemaRequestProviderType) Value() string {
	return c.value
}

func (c ListEventSchemaRequestProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventSchemaRequestProviderType) UnmarshalJSON(b []byte) error {
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
