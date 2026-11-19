package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportSubscriptionResponse Response Object
type ShowReportSubscriptionResponse struct {

	// **参数解释**: 报告ID **取值范围**: 最小值0，最大值2147483647
	ReportId *int32 `json:"report_id,omitempty"`

	// **参数解释**: 报告名称 **取值范围**: 字符长度1-32位
	ReportName *string `json:"report_name,omitempty"`

	// **参数解释**: 报告发送的时间段 **取值范围**: - morning：0点到6点 - noon：6点到12点 - afternoon：12点到18点 - evening：18点到24点
	SendingPeriod *string `json:"sending_period,omitempty"`

	// **参数解释**: 报告类别 **取值范围**: - daily_report：安全日报 - weekly_report：安全周报 - monthly_report：安全月报 - custom_report：自定义报告
	ReportCategory *string `json:"report_category,omitempty"`

	// **参数解释**: 主题urn **取值范围**: 字符长度0-256位
	TopicUrn *string `json:"topic_urn,omitempty"`

	// **参数解释**: 主题的显示名 **取值范围**: 字符长度0-256位
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**: 报告logo类型 **取值范围**: - default：华为logo - customize：自定义logo
	LogoType *string `json:"logo_type,omitempty"`

	// **参数解释**: 自定义logo的文件id **取值范围**: 字符长度0-256位
	LogoId *string `json:"logo_id,omitempty"`

	// **参数解释**: 订阅的方式 **取值范围**: - smn_topic：SMN主题 - message_center：消息中心 - silent：无需发送到邮箱
	SubscriptionType *string `json:"subscription_type,omitempty"`

	OverviewSubscriptions *OverviewSubscriptionRes `json:"overview_subscriptions,omitempty"`

	ScansSubscriptions *ScansSubscriptionRes `json:"scans_subscriptions,omitempty"`

	IntrusionSubscriptions *IntrusionSubscriptionRes `json:"intrusion_subscriptions,omitempty"`

	StatPeriod     *StatPeriodRes `json:"stat_period,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowReportSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"ShowReportSubscriptionResponse", string(data)}, " ")
}
