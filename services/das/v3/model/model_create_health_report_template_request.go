package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateHealthReportTemplateRequest Request Object
type CreateHealthReportTemplateRequest struct {

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *CreateHealthReportTemplateRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreateHealthReportTemplateReq `json:"body,omitempty"`
}

func (o CreateHealthReportTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthReportTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateHealthReportTemplateRequest", string(data)}, " ")
}

type CreateHealthReportTemplateRequestXLanguage struct {
	value string
}

type CreateHealthReportTemplateRequestXLanguageEnum struct {
	EN_US CreateHealthReportTemplateRequestXLanguage
	ZH_CN CreateHealthReportTemplateRequestXLanguage
}

func GetCreateHealthReportTemplateRequestXLanguageEnum() CreateHealthReportTemplateRequestXLanguageEnum {
	return CreateHealthReportTemplateRequestXLanguageEnum{
		EN_US: CreateHealthReportTemplateRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateHealthReportTemplateRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateHealthReportTemplateRequestXLanguage) Value() string {
	return c.value
}

func (c CreateHealthReportTemplateRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHealthReportTemplateRequestXLanguage) UnmarshalJSON(b []byte) error {
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
