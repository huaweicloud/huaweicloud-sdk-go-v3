package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityReportContentResponse Response Object
type ShowSecurityReportContentResponse struct {

	// **参数解释：** 报告ID，唯一标识当前查询的安全报告。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ReportId *string `json:"report_id,omitempty"`

	// **参数解释：** 订阅ID，关联当前报告所属的安全报告订阅记录。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubscriptionId *string `json:"subscription_id,omitempty"`

	// **参数解释：** 发送时间段，标识报告的预设发送时间（如morning表示早晨时段）。 **约束限制：** 不涉及 **取值范围：** - morning : 00:00~06:00 - noon : 06:00~12:00 - afternoon : 12:00~18:00 - evening : 12:00~18:00  **默认取值：** 不涉及
	SendingPeriod *string `json:"sending_period,omitempty"`

	// **参数解释：** 报告名称，用于标识当前安全报告的名称。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ReportName *string `json:"report_name,omitempty"`

	// **参数解释：** 报告类别，设置订阅的报告类型（如daily_report表示安全日报）。 **约束限制：** 不涉及 **取值范围：** - daily_report : 日报 - weekly_report ： 周报 - monthly_report ： 月报 - custom_report ： 自定义  **默认取值：** 不涉及
	ReportCategory *string `json:"report_category,omitempty"`

	// **参数解释：** 主题urn，关联报告发送的SMN主题唯一标识。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TopicUrn *string `json:"topic_urn,omitempty"`

	// **参数解释：** 订阅类型，接收安全报告的订阅方式（如slient表示静默订阅）。 **约束限制：** 不涉及 **取值范围：** - smn_topic : 消息主题 - slient ： 静默 - message_center ： 消息中心  **默认取值：** 不涉及
	SubscriptionType *string `json:"subscription_type,omitempty"`

	ReportContentInfo *SecurityReportContentResponseReportContentInfo `json:"report_content_info,omitempty"`

	// **参数解释：** 创建时间，报告的创建时间。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CreateTime *int64 `json:"create_time,omitempty"`

	StatPeriod     *SecurityReportContentResponseStatPeriod `json:"stat_period,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o ShowSecurityReportContentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityReportContentResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityReportContentResponse", string(data)}, " ")
}
