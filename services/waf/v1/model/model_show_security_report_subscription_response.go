package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityReportSubscriptionResponse Response Object
type ShowSecurityReportSubscriptionResponse struct {

	// **参数解释：** 订阅ID，唯一标识当前安全报告订阅。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubscriptionId *string `json:"subscription_id,omitempty"`

	// **参数解释：** 发送时间段，标识报告的预设发送时间（如morning表示早晨时段）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SendingPeriod *string `json:"sending_period,omitempty"`

	// **参数解释：** 报告名称，当前安全报告订阅的报告名称。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ReportName *string `json:"report_name,omitempty"`

	// **参数解释：** 报告类别，标识订阅的报告类型（如daily_report表示安全日报）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ReportCategory *string `json:"report_category,omitempty"`

	// **参数解释：** 主题urn，关联订阅报告发送的SMN主题唯一标识。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopicUrn *string `json:"topic_urn,omitempty"`

	// **参数解释：** 订阅类型，标识安全报告的订阅方式（如slient表示静默订阅）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubscriptionType *string `json:"subscription_type,omitempty"`

	ReportContentSubscription *SecurityReportSubscriptionResponseReportContentSubscription `json:"report_content_subscription,omitempty"`

	StatPeriod *SecurityReportSubscriptionResponseStatPeriod `json:"stat_period,omitempty"`

	// **参数解释：** 是否是所有企业项目，标识订阅是否适用于当前租户的所有企业项目（true表示适用，false表示仅适用指定项目）。 **约束限制：** 不涉及 **取值范围：** 仅支持true、false两个布尔值 **默认取值：** false
	IsAllEnterpriseProject *bool `json:"is_all_enterprise_project,omitempty"`

	// **参数解释：** 企业项目ID，订阅关联的企业项目唯一标识（is_all_enterprise_project为false时生效）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowSecurityReportSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityReportSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityReportSubscriptionResponse", string(data)}, " ")
}
