package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListRegionsRequest struct {
	// 选择接口返回的信息的语言，默认为\"zh-cn\"中文

	XLanguage *ListRegionsRequestXLanguage `json:"X-Language,omitempty"`
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
}

func GetListRegionsRequestXLanguageEnum() ListRegionsRequestXLanguageEnum {
	return ListRegionsRequestXLanguageEnum{
		ZH_CN: ListRegionsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListRegionsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListRegionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRegionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
