package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPublishedTemplatesRequest Request Object
type ListPublishedTemplatesRequest struct {

	// 语言类型，缺省值为“zh-cn”。  枚举值： - zh-cn：中文 - en-us：英文
	XLanguage *ListPublishedTemplatesRequestXLanguage `json:"X-Language,omitempty"`

	// 搜索关键字，支持按名称和描述搜索，默认null。
	Keyword *string `json:"keyword,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页的模板条数，默认10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListPublishedTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublishedTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListPublishedTemplatesRequest", string(data)}, " ")
}

type ListPublishedTemplatesRequestXLanguage struct {
	value string
}

type ListPublishedTemplatesRequestXLanguageEnum struct {
	ZH_CN ListPublishedTemplatesRequestXLanguage
	EN_US ListPublishedTemplatesRequestXLanguage
}

func GetListPublishedTemplatesRequestXLanguageEnum() ListPublishedTemplatesRequestXLanguageEnum {
	return ListPublishedTemplatesRequestXLanguageEnum{
		ZH_CN: ListPublishedTemplatesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListPublishedTemplatesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListPublishedTemplatesRequestXLanguage) Value() string {
	return c.value
}

func (c ListPublishedTemplatesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPublishedTemplatesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
