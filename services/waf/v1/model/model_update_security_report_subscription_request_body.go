package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSecurityReportSubscriptionRequestBody struct {

	// **参数解释：** 报告发送的时间段 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SendingPeriod string `json:"sending_period"`

	// **参数解释：** 报告名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ReportName string `json:"report_name"`

	// **参数解释：** 报告类别 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ReportCategory string `json:"report_category"`

	// **参数解释：** 主题urn **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopicUrn *string `json:"topic_urn,omitempty"`

	// **参数解释：** 主题的显示名 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释：** 订阅类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubscriptionType string `json:"subscription_type"`

	ReportContentSubscription *UpdateSecurityReportSubscriptionRequestBodyReportContentSubscription `json:"report_content_subscription"`

	StatPeriod *UpdateSecurityReportSubscriptionRequestBodyStatPeriod `json:"stat_period,omitempty"`

	// **参数解释：** 报告开启状态 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ReportStatus *string `json:"report_status,omitempty"`
}

func (o UpdateSecurityReportSubscriptionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityReportSubscriptionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSecurityReportSubscriptionRequestBody", string(data)}, " ")
}
