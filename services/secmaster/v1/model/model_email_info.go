package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EmailInfo 邮件信息
type EmailInfo struct {

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
}

func (o EmailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EmailInfo struct{}"
	}

	return strings.Join([]string{"EmailInfo", string(data)}, " ")
}
