package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportSlowSqlRequest Request Object
type ExportSlowSqlRequest struct {

	// 语言。
	XLanguage *ExportSlowSqlRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID。 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *ExportTableVolumeRequestBody `json:"body,omitempty"`
}

func (o ExportSlowSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSlowSqlRequest struct{}"
	}

	return strings.Join([]string{"ExportSlowSqlRequest", string(data)}, " ")
}

type ExportSlowSqlRequestXLanguage struct {
	value string
}

type ExportSlowSqlRequestXLanguageEnum struct {
	ZH_CN ExportSlowSqlRequestXLanguage
	EN_US ExportSlowSqlRequestXLanguage
}

func GetExportSlowSqlRequestXLanguageEnum() ExportSlowSqlRequestXLanguageEnum {
	return ExportSlowSqlRequestXLanguageEnum{
		ZH_CN: ExportSlowSqlRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ExportSlowSqlRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ExportSlowSqlRequestXLanguage) Value() string {
	return c.value
}

func (c ExportSlowSqlRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportSlowSqlRequestXLanguage) UnmarshalJSON(b []byte) error {
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
