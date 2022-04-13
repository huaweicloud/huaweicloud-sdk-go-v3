package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListProvidersRequest struct {
	// 分页偏移

	Offset *int32 `json:"offset,omitempty"`
	// 最大的返回数量

	Limit *int32 `json:"limit,omitempty"`
	// 选择接口返回的信息的语言，默认为\"zh-cn\"中文

	XLanguage *ListProvidersRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListProvidersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProvidersRequest struct{}"
	}

	return strings.Join([]string{"ListProvidersRequest", string(data)}, " ")
}

type ListProvidersRequestXLanguage struct {
	value string
}

type ListProvidersRequestXLanguageEnum struct {
	ZH_CN ListProvidersRequestXLanguage
	EN_US ListProvidersRequestXLanguage
}

func GetListProvidersRequestXLanguageEnum() ListProvidersRequestXLanguageEnum {
	return ListProvidersRequestXLanguageEnum{
		ZH_CN: ListProvidersRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListProvidersRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListProvidersRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProvidersRequestXLanguage) UnmarshalJSON(b []byte) error {
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
