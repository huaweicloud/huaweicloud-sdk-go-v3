package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportJobListRequest Request Object
type ExportJobListRequest struct {

	// 请求语言类型。
	XLanguage *ExportJobListRequestXLanguage `json:"X-Language,omitempty"`

	Body *ExportJobsReq `json:"body,omitempty"`
}

func (o ExportJobListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportJobListRequest struct{}"
	}

	return strings.Join([]string{"ExportJobListRequest", string(data)}, " ")
}

type ExportJobListRequestXLanguage struct {
	value string
}

type ExportJobListRequestXLanguageEnum struct {
	EN_US ExportJobListRequestXLanguage
	ZH_CN ExportJobListRequestXLanguage
	ES_US ExportJobListRequestXLanguage
	PT_BR ExportJobListRequestXLanguage
}

func GetExportJobListRequestXLanguageEnum() ExportJobListRequestXLanguageEnum {
	return ExportJobListRequestXLanguageEnum{
		EN_US: ExportJobListRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ExportJobListRequestXLanguage{
			value: "zh-cn",
		},
		ES_US: ExportJobListRequestXLanguage{
			value: "es-us",
		},
		PT_BR: ExportJobListRequestXLanguage{
			value: "pt-br",
		},
	}
}

func (c ExportJobListRequestXLanguage) Value() string {
	return c.value
}

func (c ExportJobListRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportJobListRequestXLanguage) UnmarshalJSON(b []byte) error {
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
