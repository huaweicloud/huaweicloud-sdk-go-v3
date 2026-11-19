package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddSecurityReportRequestInfo struct {

	// **参数解释**： 报告发送的时间段 **约束限制**： 不涉及 **取值范围**： 字符长度0-64位 **默认取值**： 不涉及
	SendingPeriod *string `json:"sending_period,omitempty"`

	// **参数解释**： 报告名称 **约束限制**： 必填 **取值范围**： 字符长度1-32位 **默认取值**： 不涉及
	ReportName string `json:"report_name"`

	// **参数解释**： 报告类别 **约束限制**： 不涉及 **取值范围**： 字符长度0-64位 **默认取值**： 不涉及
	ReportCategory *string `json:"report_category,omitempty"`

	// **参数解释**： 主题URN **约束限制**： 不涉及 **取值范围**： 字符长度0-256位 **默认取值**： 不涉及
	TopicUrn *string `json:"topic_urn,omitempty"`

	// **参数解释**： 主题的显示名 **约束限制**： 不涉及 **取值范围**： 字符长度0-256位 **默认取值**： 不涉及
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 订阅的方式 **约束限制**： 不涉及 **取值范围**： 字符长度0-64位 **默认取值**： 不涉及
	SubscriptionType *string `json:"subscription_type,omitempty"`

	OverviewSubscriptions *OverviewSubscriptionReq `json:"overview_subscriptions,omitempty"`

	ScansSubscriptions *ScansSubscriptionReq `json:"scans_subscriptions,omitempty"`

	IntrusionSubscriptions *IntrusionSubscriptionReq `json:"intrusion_subscriptions,omitempty"`

	StatPeriod *StatPeriodReq `json:"stat_period,omitempty"`

	// **参数解释**： 报告logo类型 **约束限制**： 不涉及 **取值范围**： - default：华为logo - customize：自定义logo  **默认取值**： 不涉及
	LogoType *string `json:"logo_type,omitempty"`

	// **参数解释**： 自定义logo的文件ID **约束限制**： 不涉及 **取值范围**： 字符长度0-256位 **默认取值**： 不涉及
	LogoId *string `json:"logo_id,omitempty"`
}

func (o AddSecurityReportRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSecurityReportRequestInfo struct{}"
	}

	return strings.Join([]string{"AddSecurityReportRequestInfo", string(data)}, " ")
}
