package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportRuleInfo 报告发送规则
type ReportRuleInfo struct {

	// 报告发送规则id
	Id *string `json:"id,omitempty"`

	// 租户ID
	ProjectId *string `json:"project_id,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 数据周期
	Cycle string `json:"cycle"`

	// cron表达式
	Rule string `json:"rule"`

	// 报告数据周期开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 报告数据周期结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 邮件标题
	EmailTitle *string `json:"email_title,omitempty"`

	// 收件人邮箱
	EmailTo *string `json:"email_to,omitempty"`

	// 邮件内容
	EmailContent *string `json:"email_content,omitempty"`

	// 报告类型，支持图片、pdf、html
	ReportFileType *string `json:"report_file_type,omitempty"`
}

func (o ReportRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportRuleInfo struct{}"
	}

	return strings.Join([]string{"ReportRuleInfo", string(data)}, " ")
}
