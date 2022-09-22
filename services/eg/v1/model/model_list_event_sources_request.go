package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListEventSourcesRequest struct {

	// 指定查询的事件通道ID
	ChannelId *string `json:"channel_id,omitempty"`

	// 偏移量，表示从此偏移量开始查询，偏移量不能小于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，不能小于0
	Limit *int32 `json:"limit,omitempty"`

	// 指定查询排序
	Sort *string `json:"sort,omitempty"`

	// 指定查询提供方的类型
	ProviderType *ListEventSourcesRequestProviderType `json:"provider_type,omitempty"`

	// 指定查询的事件源名称，精准匹配
	Name *string `json:"name,omitempty"`

	// 指定查询的事件源名称，模糊匹配
	FuzzyName *string `json:"fuzzy_name,omitempty"`

	// 指定查询的事件源标签，模糊匹配
	FuzzyLabel *string `json:"fuzzy_label,omitempty"`
}

func (o ListEventSourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventSourcesRequest struct{}"
	}

	return strings.Join([]string{"ListEventSourcesRequest", string(data)}, " ")
}

type ListEventSourcesRequestProviderType struct {
	value string
}

type ListEventSourcesRequestProviderTypeEnum struct {
	OFFICIAL ListEventSourcesRequestProviderType
	CUSTOM   ListEventSourcesRequestProviderType
}

func GetListEventSourcesRequestProviderTypeEnum() ListEventSourcesRequestProviderTypeEnum {
	return ListEventSourcesRequestProviderTypeEnum{
		OFFICIAL: ListEventSourcesRequestProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: ListEventSourcesRequestProviderType{
			value: "CUSTOM",
		},
	}
}

func (c ListEventSourcesRequestProviderType) Value() string {
	return c.value
}

func (c ListEventSourcesRequestProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventSourcesRequestProviderType) UnmarshalJSON(b []byte) error {
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
