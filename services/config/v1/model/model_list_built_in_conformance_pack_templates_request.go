package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBuiltInConformancePackTemplatesRequest Request Object
type ListBuiltInConformancePackTemplatesRequest struct {

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`

	// 预定义合规包模板名称。
	TemplateKey *string `json:"template_key,omitempty"`

	// 选择接口返回的信息的语言，默认为\"zh-cn\"中文
	XLanguage *ListBuiltInConformancePackTemplatesRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListBuiltInConformancePackTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBuiltInConformancePackTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListBuiltInConformancePackTemplatesRequest", string(data)}, " ")
}

type ListBuiltInConformancePackTemplatesRequestXLanguage struct {
	value string
}

type ListBuiltInConformancePackTemplatesRequestXLanguageEnum struct {
	ZH_CN ListBuiltInConformancePackTemplatesRequestXLanguage
	EN_US ListBuiltInConformancePackTemplatesRequestXLanguage
}

func GetListBuiltInConformancePackTemplatesRequestXLanguageEnum() ListBuiltInConformancePackTemplatesRequestXLanguageEnum {
	return ListBuiltInConformancePackTemplatesRequestXLanguageEnum{
		ZH_CN: ListBuiltInConformancePackTemplatesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListBuiltInConformancePackTemplatesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListBuiltInConformancePackTemplatesRequestXLanguage) Value() string {
	return c.value
}

func (c ListBuiltInConformancePackTemplatesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBuiltInConformancePackTemplatesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
