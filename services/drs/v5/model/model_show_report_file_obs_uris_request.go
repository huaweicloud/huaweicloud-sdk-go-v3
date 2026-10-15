package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowReportFileObsUrisRequest Request Object
type ShowReportFileObsUrisRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 导出的sql文件类型。取值范围：  abnormal_sql ：异常sql列表 abnormal_sql_detail ：异常sql详情 slow_sql ：慢sql列表 slow_sql_detail ： 慢sql详情
	ReportType string `json:"report_type"`

	// 请求语言类型。
	XLanguage *ShowReportFileObsUrisRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowReportFileObsUrisRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportFileObsUrisRequest struct{}"
	}

	return strings.Join([]string{"ShowReportFileObsUrisRequest", string(data)}, " ")
}

type ShowReportFileObsUrisRequestXLanguage struct {
	value string
}

type ShowReportFileObsUrisRequestXLanguageEnum struct {
	EN_US ShowReportFileObsUrisRequestXLanguage
	ZH_CN ShowReportFileObsUrisRequestXLanguage
}

func GetShowReportFileObsUrisRequestXLanguageEnum() ShowReportFileObsUrisRequestXLanguageEnum {
	return ShowReportFileObsUrisRequestXLanguageEnum{
		EN_US: ShowReportFileObsUrisRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowReportFileObsUrisRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowReportFileObsUrisRequestXLanguage) Value() string {
	return c.value
}

func (c ShowReportFileObsUrisRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowReportFileObsUrisRequestXLanguage) UnmarshalJSON(b []byte) error {
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
