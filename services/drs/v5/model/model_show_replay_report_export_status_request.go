package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowReplayReportExportStatusRequest Request Object
type ShowReplayReportExportStatusRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowReplayReportExportStatusRequestXLanguage `json:"X-Language,omitempty"`

	// 导出的sql文件类型。取值范围： - abnormal_sql ：异常sql列表 - abnormal_sql_detail ：异常sql详情 - slow_sql ：慢sql列表 - slow_sql_detail ： 慢sql详情
	FileType string `json:"file_type"`
}

func (o ShowReplayReportExportStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplayReportExportStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowReplayReportExportStatusRequest", string(data)}, " ")
}

type ShowReplayReportExportStatusRequestXLanguage struct {
	value string
}

type ShowReplayReportExportStatusRequestXLanguageEnum struct {
	EN_US ShowReplayReportExportStatusRequestXLanguage
	ZH_CN ShowReplayReportExportStatusRequestXLanguage
}

func GetShowReplayReportExportStatusRequestXLanguageEnum() ShowReplayReportExportStatusRequestXLanguageEnum {
	return ShowReplayReportExportStatusRequestXLanguageEnum{
		EN_US: ShowReplayReportExportStatusRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowReplayReportExportStatusRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowReplayReportExportStatusRequestXLanguage) Value() string {
	return c.value
}

func (c ShowReplayReportExportStatusRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowReplayReportExportStatusRequestXLanguage) UnmarshalJSON(b []byte) error {
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
