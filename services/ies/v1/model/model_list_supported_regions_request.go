package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListSupportedRegionsRequest struct {

	// 选择接口返回信息的语言类型，默认为中文\"zh-cn\"
	XLanguage *ListSupportedRegionsRequestXLanguage `json:"X-Language,omitempty"`

	// 每页的数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页标识
	Marker *string `json:"marker,omitempty"`

	// 排序字段
	SortKey *[]string `json:"sort_key,omitempty"`

	// 排序方向，取值范围： - desc：降序 - acs：升序
	SortDir *[]string `json:"sort_dir,omitempty"`
}

func (o ListSupportedRegionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportedRegionsRequest struct{}"
	}

	return strings.Join([]string{"ListSupportedRegionsRequest", string(data)}, " ")
}

type ListSupportedRegionsRequestXLanguage struct {
	value string
}

type ListSupportedRegionsRequestXLanguageEnum struct {
	ZH_CN ListSupportedRegionsRequestXLanguage
	EN_US ListSupportedRegionsRequestXLanguage
	PT_BR ListSupportedRegionsRequestXLanguage
	ES_US ListSupportedRegionsRequestXLanguage
	ES_ES ListSupportedRegionsRequestXLanguage
}

func GetListSupportedRegionsRequestXLanguageEnum() ListSupportedRegionsRequestXLanguageEnum {
	return ListSupportedRegionsRequestXLanguageEnum{
		ZH_CN: ListSupportedRegionsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSupportedRegionsRequestXLanguage{
			value: "en-us",
		},
		PT_BR: ListSupportedRegionsRequestXLanguage{
			value: "pt-br",
		},
		ES_US: ListSupportedRegionsRequestXLanguage{
			value: "es-us",
		},
		ES_ES: ListSupportedRegionsRequestXLanguage{
			value: "es-es",
		},
	}
}

func (c ListSupportedRegionsRequestXLanguage) Value() string {
	return c.value
}

func (c ListSupportedRegionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupportedRegionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
