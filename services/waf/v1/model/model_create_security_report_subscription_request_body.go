package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSecurityReportSubscriptionRequestBody struct {

	// **参数解释：** 发送时间段，设置报告的预设发送时间（如morning表示早晨时段）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SendingPeriod *string `json:"sending_period,omitempty"`

	// **参数解释：** 报告名称，设置订阅报告的名称。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ReportName *string `json:"report_name,omitempty"`

	// **参数解释：** 报告类别，设置订阅的报告类型（如daily_report表示安全日报）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ReportCategory *string `json:"report_category,omitempty"`

	// **参数解释：** 主题urn，设置报告发送的SMN主题唯一标识。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopicUrn *string `json:"topic_urn,omitempty"`

	// **参数解释：** 订阅类型，设置报告的订阅方式（如smn_topic表示SMN主题订阅）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubscriptionType *string `json:"subscription_type,omitempty"`

	ReportContentSubscription *CreateSecurityReportSubscriptionRequestBodyReportContentSubscription `json:"report_content_subscription,omitempty"`

	StatPeriod *CreateSecurityReportSubscriptionRequestBodyStatPeriod `json:"stat_period,omitempty"`

	// **参数解释：** 控制台URL前缀，用于拼接生成报告中相关资源的控制台访问链接 **格式要求：** 必须以http://或https://开头的有效URL格式 **约束限制：** 仅支持华为云控制台域名及路径 **默认取值：** https://console.huaweicloud.com/console/
	ConsoleUrlPrefix *string `json:"console_url_prefix,omitempty"`
}

func (o CreateSecurityReportSubscriptionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityReportSubscriptionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecurityReportSubscriptionRequestBody", string(data)}, " ")
}
