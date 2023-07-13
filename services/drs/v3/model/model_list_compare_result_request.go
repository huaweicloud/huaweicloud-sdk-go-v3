package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCompareResultRequest Request Object
type ListCompareResultRequest struct {

	// 请求语言类型
	XLanguage *ListCompareResultRequestXLanguage `json:"X-Language,omitempty"`

	Body *QueryCompareResultReq `json:"body,omitempty"`
}

func (o ListCompareResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCompareResultRequest struct{}"
	}

	return strings.Join([]string{"ListCompareResultRequest", string(data)}, " ")
}

type ListCompareResultRequestXLanguage struct {
	value string
}

type ListCompareResultRequestXLanguageEnum struct {
	EN_US ListCompareResultRequestXLanguage
	ZH_CN ListCompareResultRequestXLanguage
}

func GetListCompareResultRequestXLanguageEnum() ListCompareResultRequestXLanguageEnum {
	return ListCompareResultRequestXLanguageEnum{
		EN_US: ListCompareResultRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListCompareResultRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListCompareResultRequestXLanguage) Value() string {
	return c.value
}

func (c ListCompareResultRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCompareResultRequestXLanguage) UnmarshalJSON(b []byte) error {
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
