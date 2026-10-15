package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateHealthReportTemplateRequest Request Object
type UpdateHealthReportTemplateRequest struct {

	// 健康报告模板ID。
	TemplateId string `json:"template_id"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *UpdateHealthReportTemplateRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpdateHealthReportTemplateReq `json:"body,omitempty"`
}

func (o UpdateHealthReportTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthReportTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateHealthReportTemplateRequest", string(data)}, " ")
}

type UpdateHealthReportTemplateRequestXLanguage struct {
	value string
}

type UpdateHealthReportTemplateRequestXLanguageEnum struct {
	EN_US UpdateHealthReportTemplateRequestXLanguage
	ZH_CN UpdateHealthReportTemplateRequestXLanguage
}

func GetUpdateHealthReportTemplateRequestXLanguageEnum() UpdateHealthReportTemplateRequestXLanguageEnum {
	return UpdateHealthReportTemplateRequestXLanguageEnum{
		EN_US: UpdateHealthReportTemplateRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UpdateHealthReportTemplateRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UpdateHealthReportTemplateRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateHealthReportTemplateRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHealthReportTemplateRequestXLanguage) UnmarshalJSON(b []byte) error {
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
