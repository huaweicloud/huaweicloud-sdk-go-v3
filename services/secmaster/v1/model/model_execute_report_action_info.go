package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteReportActionInfo 下载安全报告的请求
type ExecuteReportActionInfo struct {

	// 报告的base64编码
	ReportContent *string `json:"report_content,omitempty"`

	// 邮件标题
	EmailTitle *string `json:"email_title,omitempty"`

	// 收件人邮箱
	EmailTo *string `json:"email_to,omitempty"`

	// 抄送人邮箱
	EmailCc *string `json:"email_cc,omitempty"`

	// 邮件内容
	EmailContent *string `json:"email_content,omitempty"`

	// 附件类型
	ReportFileType *string `json:"report_file_type,omitempty"`

	// API的动作支持send和download
	Action ExecuteReportActionInfoAction `json:"action"`

	// 安全报告Base64编码的内容
	ReportPage *string `json:"report_page,omitempty"`

	PageConfig *PageConfigPx `json:"page_config,omitempty"`
}

func (o ExecuteReportActionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteReportActionInfo struct{}"
	}

	return strings.Join([]string{"ExecuteReportActionInfo", string(data)}, " ")
}

type ExecuteReportActionInfoAction struct {
	value string
}

type ExecuteReportActionInfoActionEnum struct {
	SEND     ExecuteReportActionInfoAction
	DOWNLOAD ExecuteReportActionInfoAction
}

func GetExecuteReportActionInfoActionEnum() ExecuteReportActionInfoActionEnum {
	return ExecuteReportActionInfoActionEnum{
		SEND: ExecuteReportActionInfoAction{
			value: "send",
		},
		DOWNLOAD: ExecuteReportActionInfoAction{
			value: "download",
		},
	}
}

func (c ExecuteReportActionInfoAction) Value() string {
	return c.value
}

func (c ExecuteReportActionInfoAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteReportActionInfoAction) UnmarshalJSON(b []byte) error {
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
