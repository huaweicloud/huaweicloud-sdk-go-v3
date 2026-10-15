package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteHealthReportTemplateRequest Request Object
type DeleteHealthReportTemplateRequest struct {

	// 健康报告模板ID。
	TemplateId string `json:"template_id"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *DeleteHealthReportTemplateRequestXLanguage `json:"X-Language,omitempty"`
}

func (o DeleteHealthReportTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHealthReportTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteHealthReportTemplateRequest", string(data)}, " ")
}

type DeleteHealthReportTemplateRequestXLanguage struct {
	value string
}

type DeleteHealthReportTemplateRequestXLanguageEnum struct {
	EN_US DeleteHealthReportTemplateRequestXLanguage
	ZH_CN DeleteHealthReportTemplateRequestXLanguage
}

func GetDeleteHealthReportTemplateRequestXLanguageEnum() DeleteHealthReportTemplateRequestXLanguageEnum {
	return DeleteHealthReportTemplateRequestXLanguageEnum{
		EN_US: DeleteHealthReportTemplateRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DeleteHealthReportTemplateRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c DeleteHealthReportTemplateRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteHealthReportTemplateRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteHealthReportTemplateRequestXLanguage) UnmarshalJSON(b []byte) error {
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
