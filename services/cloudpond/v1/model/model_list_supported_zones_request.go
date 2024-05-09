package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSupportedZonesRequest Request Object
type ListSupportedZonesRequest struct {

	// 选择接口返回信息的语言类型，默认为中文\"zh-cn\"
	XLanguage *ListSupportedZonesRequestXLanguage `json:"X-Language,omitempty"`

	// 每页的数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页标识
	Marker *string `json:"marker,omitempty"`

	// 排序字段
	SortKey *[]string `json:"sort_key,omitempty"`

	// 排序方向，取值范围： - desc：降序 - acs：升序
	SortDir *[]string `json:"sort_dir,omitempty"`
}

func (o ListSupportedZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportedZonesRequest struct{}"
	}

	return strings.Join([]string{"ListSupportedZonesRequest", string(data)}, " ")
}

type ListSupportedZonesRequestXLanguage struct {
	value string
}

type ListSupportedZonesRequestXLanguageEnum struct {
	ZH_CN ListSupportedZonesRequestXLanguage
	EN_US ListSupportedZonesRequestXLanguage
	PT_BR ListSupportedZonesRequestXLanguage
	ES_US ListSupportedZonesRequestXLanguage
	ES_ES ListSupportedZonesRequestXLanguage
}

func GetListSupportedZonesRequestXLanguageEnum() ListSupportedZonesRequestXLanguageEnum {
	return ListSupportedZonesRequestXLanguageEnum{
		ZH_CN: ListSupportedZonesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSupportedZonesRequestXLanguage{
			value: "en-us",
		},
		PT_BR: ListSupportedZonesRequestXLanguage{
			value: "pt-br",
		},
		ES_US: ListSupportedZonesRequestXLanguage{
			value: "es-us",
		},
		ES_ES: ListSupportedZonesRequestXLanguage{
			value: "es-es",
		},
	}
}

func (c ListSupportedZonesRequestXLanguage) Value() string {
	return c.value
}

func (c ListSupportedZonesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupportedZonesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
