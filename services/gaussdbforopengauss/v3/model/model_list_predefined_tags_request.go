package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPredefinedTagsRequest Request Object
type ListPredefinedTagsRequest struct {

	// 语言
	XLanguage *ListPredefinedTagsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListPredefinedTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPredefinedTagsRequest struct{}"
	}

	return strings.Join([]string{"ListPredefinedTagsRequest", string(data)}, " ")
}

type ListPredefinedTagsRequestXLanguage struct {
	value string
}

type ListPredefinedTagsRequestXLanguageEnum struct {
	ZH_CN ListPredefinedTagsRequestXLanguage
	EN_US ListPredefinedTagsRequestXLanguage
}

func GetListPredefinedTagsRequestXLanguageEnum() ListPredefinedTagsRequestXLanguageEnum {
	return ListPredefinedTagsRequestXLanguageEnum{
		ZH_CN: ListPredefinedTagsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListPredefinedTagsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListPredefinedTagsRequestXLanguage) Value() string {
	return c.value
}

func (c ListPredefinedTagsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPredefinedTagsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
