package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListHealthReportTemplateRequest Request Object
type ListHealthReportTemplateRequest struct {

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *ListHealthReportTemplateRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListHealthReportTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthReportTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListHealthReportTemplateRequest", string(data)}, " ")
}

type ListHealthReportTemplateRequestXLanguage struct {
	value string
}

type ListHealthReportTemplateRequestXLanguageEnum struct {
	EN_US ListHealthReportTemplateRequestXLanguage
	ZH_CN ListHealthReportTemplateRequestXLanguage
}

func GetListHealthReportTemplateRequestXLanguageEnum() ListHealthReportTemplateRequestXLanguageEnum {
	return ListHealthReportTemplateRequestXLanguageEnum{
		EN_US: ListHealthReportTemplateRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListHealthReportTemplateRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListHealthReportTemplateRequestXLanguage) Value() string {
	return c.value
}

func (c ListHealthReportTemplateRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHealthReportTemplateRequestXLanguage) UnmarshalJSON(b []byte) error {
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
