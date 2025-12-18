package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportSlowSqlListRequest Request Object
type ExportSlowSqlListRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ExportSlowSqlListRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *ListSlowSqlsRequestBody `json:"body,omitempty"`
}

func (o ExportSlowSqlListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSlowSqlListRequest struct{}"
	}

	return strings.Join([]string{"ExportSlowSqlListRequest", string(data)}, " ")
}

type ExportSlowSqlListRequestXLanguage struct {
	value string
}

type ExportSlowSqlListRequestXLanguageEnum struct {
	ZH_CN ExportSlowSqlListRequestXLanguage
	EN_US ExportSlowSqlListRequestXLanguage
}

func GetExportSlowSqlListRequestXLanguageEnum() ExportSlowSqlListRequestXLanguageEnum {
	return ExportSlowSqlListRequestXLanguageEnum{
		ZH_CN: ExportSlowSqlListRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ExportSlowSqlListRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ExportSlowSqlListRequestXLanguage) Value() string {
	return c.value
}

func (c ExportSlowSqlListRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportSlowSqlListRequestXLanguage) UnmarshalJSON(b []byte) error {
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
