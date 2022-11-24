package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListChannelsRequest struct {

	// 偏移量，表示从此偏移量开始查询，偏移量不能小于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，不能小于0
	Limit *int32 `json:"limit,omitempty"`

	// 指定查询排序
	Sort *string `json:"sort,omitempty"`

	// 指定查询提供方的类型
	ProviderType *ListChannelsRequestProviderType `json:"provider_type,omitempty"`

	// 指定查询的事件通道名称，精准匹配
	Name *string `json:"name,omitempty"`

	// 指定查询的事件通道名称，模糊匹配
	FuzzyName *string `json:"fuzzy_name,omitempty"`
}

func (o ListChannelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListChannelsRequest struct{}"
	}

	return strings.Join([]string{"ListChannelsRequest", string(data)}, " ")
}

type ListChannelsRequestProviderType struct {
	value string
}

type ListChannelsRequestProviderTypeEnum struct {
	OFFICIAL ListChannelsRequestProviderType
	CUSTOM   ListChannelsRequestProviderType
}

func GetListChannelsRequestProviderTypeEnum() ListChannelsRequestProviderTypeEnum {
	return ListChannelsRequestProviderTypeEnum{
		OFFICIAL: ListChannelsRequestProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: ListChannelsRequestProviderType{
			value: "CUSTOM",
		},
	}
}

func (c ListChannelsRequestProviderType) Value() string {
	return c.value
}

func (c ListChannelsRequestProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListChannelsRequestProviderType) UnmarshalJSON(b []byte) error {
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
