package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowHealthReportTemplateRequest Request Object
type ShowHealthReportTemplateRequest struct {

	// 健康报告模板ID。
	TemplateId string `json:"template_id"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *ShowHealthReportTemplateRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowHealthReportTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthReportTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowHealthReportTemplateRequest", string(data)}, " ")
}

type ShowHealthReportTemplateRequestXLanguage struct {
	value string
}

type ShowHealthReportTemplateRequestXLanguageEnum struct {
	EN_US ShowHealthReportTemplateRequestXLanguage
	ZH_CN ShowHealthReportTemplateRequestXLanguage
}

func GetShowHealthReportTemplateRequestXLanguageEnum() ShowHealthReportTemplateRequestXLanguageEnum {
	return ShowHealthReportTemplateRequestXLanguageEnum{
		EN_US: ShowHealthReportTemplateRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowHealthReportTemplateRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowHealthReportTemplateRequestXLanguage) Value() string {
	return c.value
}

func (c ShowHealthReportTemplateRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHealthReportTemplateRequestXLanguage) UnmarshalJSON(b []byte) error {
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
