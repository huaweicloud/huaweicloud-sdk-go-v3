package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityReportSubscriptionSummaryResponse struct {

	// **参数解释：** 订阅ID **取值范围：** 不涉及
	SubscriptionId *string `json:"subscription_id,omitempty"`

	// **参数解释：** 报告ID **取值范围：** 不涉及
	ReportId *string `json:"report_id,omitempty"`

	// **参数解释：** 报告模板名称 **取值范围：** 只能由中文、字母、数字和括号内所列符号（_-.:：）组成，且不能超过256个字符长度。
	ReportName *string `json:"report_name,omitempty"`

	// **参数解释：** 发送时间段 **取值范围：** - morning：00:00~06:00 - noon：06:00~12:00 - afternoon：12:00~18:00 - evening：18:00~24:00
	SendingPeriod *string `json:"sending_period,omitempty"`

	// **参数解释：** 报告类型 **取值范围：** - daily_report：安全日报 - weekly_report：安全周报 - monthly_report：安全月报 - custom_report：自定义报告
	ReportCategory *string `json:"report_category,omitempty"`

	// **参数解释：** 开启状态 **取值范围：** - opened：已开启 - closed：已关闭
	ReportStatus *string `json:"report_status,omitempty"`

	// **参数解释：** 是否是所有企业项目 **取值范围：** - true：是所有企业项目 - false：不是所有企业项目
	IsAllEnterpriseProject *bool `json:"is_all_enterprise_project,omitempty"`

	// **参数解释：** 企业项目ID，当前报告统计的企业项目范围 **取值范围：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 报告所属项目，生成的报告所归属的企业项目 **取值范围：** 不涉及
	TemplateEpsId *string `json:"template_eps_id,omitempty"`

	// **参数解释：** 报告生成状态 **取值范围：** - true：已生成 - false：暂未生成
	IsReportCreated *bool `json:"is_report_created,omitempty"`

	// **参数解释：** 报告生成时间 **取值范围：** 13位毫秒时间戳
	LatestCreateTime *int64 `json:"latest_create_time,omitempty"`
}

func (o SecurityReportSubscriptionSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportSubscriptionSummaryResponse struct{}"
	}

	return strings.Join([]string{"SecurityReportSubscriptionSummaryResponse", string(data)}, " ")
}
