package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeSecurityReportRequestInfo 修改报告
type ChangeSecurityReportRequestInfo struct {

	// **参数解释**: 报告ID **约束限制**: 必填 **取值范围**: 最小值0，最大值2147483647 **默认取值**: 不涉及
	ReportId int32 `json:"report_id"`

	// **参数解释**: 报告发送的时间段 **约束限制**: 不涉及 **取值范围**: - morning：0点到6点 - noon：6点到12点 - afternoon：12点到18点 - evening：18点到24点  **默认取值**: 不涉及
	SendingPeriod *string `json:"sending_period,omitempty"`

	// **参数解释**: 报告名称 **约束限制**: 必填 **取值范围**: 字符长度1-32位 **默认取值**: 不涉及
	ReportName string `json:"report_name"`

	// **参数解释**: 主题urn **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	TopicUrn *string `json:"topic_urn,omitempty"`

	// **参数解释**: 主题的显示名 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**: 订阅的方式 **约束限制**: 不涉及 **取值范围**: - smn_topic：SMN主题 - message_center：消息中心 - silent：无需发送到邮箱  **默认取值**: 不涉及
	SubscriptionType *string `json:"subscription_type,omitempty"`

	OverviewSubscriptions *OverviewSubscriptionReq `json:"overview_subscriptions,omitempty"`

	ScansSubscriptions *ScansSubscriptionReq `json:"scans_subscriptions,omitempty"`

	IntrusionSubscriptions *IntrusionSubscriptionReq `json:"intrusion_subscriptions,omitempty"`

	// **参数解释**: 报告logo类型 **约束限制**: 不涉及 **取值范围**: - default：华为logo - customize：自定义logo  **默认取值**: 不涉及
	LogoType *string `json:"logo_type,omitempty"`

	// **参数解释**: 自定义logo的文件id **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	LogoId *string `json:"logo_id,omitempty"`
}

func (o ChangeSecurityReportRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityReportRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeSecurityReportRequestInfo", string(data)}, " ")
}
