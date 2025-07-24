package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportCreationTemplateRequest Request Object
type ExportCreationTemplateRequest struct {

	// 请求语言类型。
	XLanguage *ExportCreationTemplateRequestXLanguage `json:"X-Language,omitempty"`

	Body *ExportJobsTemplateReq `json:"body,omitempty"`
}

func (o ExportCreationTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCreationTemplateRequest struct{}"
	}

	return strings.Join([]string{"ExportCreationTemplateRequest", string(data)}, " ")
}

type ExportCreationTemplateRequestXLanguage struct {
	value string
}

type ExportCreationTemplateRequestXLanguageEnum struct {
	EN_US ExportCreationTemplateRequestXLanguage
	ZH_CN ExportCreationTemplateRequestXLanguage
}

func GetExportCreationTemplateRequestXLanguageEnum() ExportCreationTemplateRequestXLanguageEnum {
	return ExportCreationTemplateRequestXLanguageEnum{
		EN_US: ExportCreationTemplateRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ExportCreationTemplateRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ExportCreationTemplateRequestXLanguage) Value() string {
	return c.value
}

func (c ExportCreationTemplateRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportCreationTemplateRequestXLanguage) UnmarshalJSON(b []byte) error {
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
