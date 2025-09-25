package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAllProvidersRequest Request Object
type ListAllProvidersRequest struct {

	// 分页偏移
	Offset *int32 `json:"offset,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 选择接口返回的信息的语言，默认为\"zh-cn\"中文
	XLanguage *ListAllProvidersRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListAllProvidersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllProvidersRequest struct{}"
	}

	return strings.Join([]string{"ListAllProvidersRequest", string(data)}, " ")
}

type ListAllProvidersRequestXLanguage struct {
	value string
}

type ListAllProvidersRequestXLanguageEnum struct {
	ZH_CN ListAllProvidersRequestXLanguage
	EN_US ListAllProvidersRequestXLanguage
	FR_FR ListAllProvidersRequestXLanguage
	ES_US ListAllProvidersRequestXLanguage
	PT_BR ListAllProvidersRequestXLanguage
}

func GetListAllProvidersRequestXLanguageEnum() ListAllProvidersRequestXLanguageEnum {
	return ListAllProvidersRequestXLanguageEnum{
		ZH_CN: ListAllProvidersRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListAllProvidersRequestXLanguage{
			value: "en-us",
		},
		FR_FR: ListAllProvidersRequestXLanguage{
			value: "fr-fr",
		},
		ES_US: ListAllProvidersRequestXLanguage{
			value: "es-us",
		},
		PT_BR: ListAllProvidersRequestXLanguage{
			value: "pt-br",
		},
	}
}

func (c ListAllProvidersRequestXLanguage) Value() string {
	return c.value
}

func (c ListAllProvidersRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAllProvidersRequestXLanguage) UnmarshalJSON(b []byte) error {
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
