package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSecurityReportSendingRecordResponseItems struct {

	// **参数解释：** 报告ID，唯一标识该发送记录对应的安全报告。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ReportId string `json:"report_id"`

	// **参数解释：** 订阅ID，关联该发送记录所属的安全报告订阅。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubscriptionId string `json:"subscription_id"`

	// **参数解释：** 报告名称，该发送记录对应的安全报告名称。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ReportName string `json:"report_name"`

	StatPeriod *ListSecurityReportSendingRecordResponseStatPeriod `json:"stat_period"`

	// **参数解释：** 报告类别，设置订阅的报告类型（如daily_report表示安全日报）。 **约束限制：** 不涉及 **取值范围：** - daily_report : 日报 - weekly_report ： 周报 - monthly_report ： 月报 - custom_report ： 自定义  **默认取值：** 不涉及
	ReportCategory string `json:"report_category"`

	// **参数解释：** 发送时间，该报告实际发送的时间戳（毫秒级）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SendingTime int64 `json:"sending_time"`
}

func (o ListSecurityReportSendingRecordResponseItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityReportSendingRecordResponseItems struct{}"
	}

	return strings.Join([]string{"ListSecurityReportSendingRecordResponseItems", string(data)}, " ")
}
