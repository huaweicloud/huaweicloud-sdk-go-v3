package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportReplayReportRequest Request Object
type ExportReplayReportRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ExportReplayReportRequestXLanguage `json:"X-Language,omitempty"`

	Body *ExportSqlDataReq `json:"body,omitempty"`
}

func (o ExportReplayReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportReplayReportRequest struct{}"
	}

	return strings.Join([]string{"ExportReplayReportRequest", string(data)}, " ")
}

type ExportReplayReportRequestXLanguage struct {
	value string
}

type ExportReplayReportRequestXLanguageEnum struct {
	EN_US ExportReplayReportRequestXLanguage
	ZH_CN ExportReplayReportRequestXLanguage
}

func GetExportReplayReportRequestXLanguageEnum() ExportReplayReportRequestXLanguageEnum {
	return ExportReplayReportRequestXLanguageEnum{
		EN_US: ExportReplayReportRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ExportReplayReportRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ExportReplayReportRequestXLanguage) Value() string {
	return c.value
}

func (c ExportReplayReportRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportReplayReportRequestXLanguage) UnmarshalJSON(b []byte) error {
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
