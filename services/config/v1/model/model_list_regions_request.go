package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRegionsRequest Request Object
type ListRegionsRequest struct {

	// 选择接口返回的信息的语言，默认为\"zh-cn\"中文
	XLanguage *ListRegionsRequestXLanguage `json:"X-Language,omitempty"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`
}

func (o ListRegionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegionsRequest struct{}"
	}

	return strings.Join([]string{"ListRegionsRequest", string(data)}, " ")
}

type ListRegionsRequestXLanguage struct {
	value string
}

type ListRegionsRequestXLanguageEnum struct {
	ZH_CN ListRegionsRequestXLanguage
	EN_US ListRegionsRequestXLanguage
	FR_FR ListRegionsRequestXLanguage
	ES_US ListRegionsRequestXLanguage
	PT_BR ListRegionsRequestXLanguage
}

func GetListRegionsRequestXLanguageEnum() ListRegionsRequestXLanguageEnum {
	return ListRegionsRequestXLanguageEnum{
		ZH_CN: ListRegionsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListRegionsRequestXLanguage{
			value: "en-us",
		},
		FR_FR: ListRegionsRequestXLanguage{
			value: "fr-fr",
		},
		ES_US: ListRegionsRequestXLanguage{
			value: "es-us",
		},
		PT_BR: ListRegionsRequestXLanguage{
			value: "pt-br",
		},
	}
}

func (c ListRegionsRequestXLanguage) Value() string {
	return c.value
}

func (c ListRegionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRegionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
