package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListHealthReportConfigsRequest Request Object
type ListHealthReportConfigsRequest struct {

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *ListHealthReportConfigsRequestXLanguage `json:"X-Language,omitempty"`

	Body *ListHealthReportConfigsReq `json:"body,omitempty"`
}

func (o ListHealthReportConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthReportConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListHealthReportConfigsRequest", string(data)}, " ")
}

type ListHealthReportConfigsRequestXLanguage struct {
	value string
}

type ListHealthReportConfigsRequestXLanguageEnum struct {
	EN_US ListHealthReportConfigsRequestXLanguage
	ZH_CN ListHealthReportConfigsRequestXLanguage
}

func GetListHealthReportConfigsRequestXLanguageEnum() ListHealthReportConfigsRequestXLanguageEnum {
	return ListHealthReportConfigsRequestXLanguageEnum{
		EN_US: ListHealthReportConfigsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListHealthReportConfigsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListHealthReportConfigsRequestXLanguage) Value() string {
	return c.value
}

func (c ListHealthReportConfigsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHealthReportConfigsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
