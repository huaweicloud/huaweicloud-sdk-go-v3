package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportSlowSqlStatisticsRequest Request Object
type ExportSlowSqlStatisticsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。
	XLanguage *ExportSlowSqlStatisticsRequestXLanguage `json:"X-Language,omitempty"`

	Body *ExportSlowSqlStatisticsRequestBody `json:"body,omitempty"`
}

func (o ExportSlowSqlStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSlowSqlStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ExportSlowSqlStatisticsRequest", string(data)}, " ")
}

type ExportSlowSqlStatisticsRequestXLanguage struct {
	value string
}

type ExportSlowSqlStatisticsRequestXLanguageEnum struct {
	EN_US ExportSlowSqlStatisticsRequestXLanguage
	ZH_CN ExportSlowSqlStatisticsRequestXLanguage
}

func GetExportSlowSqlStatisticsRequestXLanguageEnum() ExportSlowSqlStatisticsRequestXLanguageEnum {
	return ExportSlowSqlStatisticsRequestXLanguageEnum{
		EN_US: ExportSlowSqlStatisticsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ExportSlowSqlStatisticsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ExportSlowSqlStatisticsRequestXLanguage) Value() string {
	return c.value
}

func (c ExportSlowSqlStatisticsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportSlowSqlStatisticsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
