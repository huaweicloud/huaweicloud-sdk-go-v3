package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportInstanceInfosRequest Request Object
type ExportInstanceInfosRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ExportInstanceInfosRequestXLanguage `json:"X-Language,omitempty"`

	Body *ExportInstanceInfosRequestBody `json:"body,omitempty"`
}

func (o ExportInstanceInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportInstanceInfosRequest struct{}"
	}

	return strings.Join([]string{"ExportInstanceInfosRequest", string(data)}, " ")
}

type ExportInstanceInfosRequestXLanguage struct {
	value string
}

type ExportInstanceInfosRequestXLanguageEnum struct {
	ZH_CN ExportInstanceInfosRequestXLanguage
	EN_US ExportInstanceInfosRequestXLanguage
}

func GetExportInstanceInfosRequestXLanguageEnum() ExportInstanceInfosRequestXLanguageEnum {
	return ExportInstanceInfosRequestXLanguageEnum{
		ZH_CN: ExportInstanceInfosRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ExportInstanceInfosRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ExportInstanceInfosRequestXLanguage) Value() string {
	return c.value
}

func (c ExportInstanceInfosRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportInstanceInfosRequestXLanguage) UnmarshalJSON(b []byte) error {
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
