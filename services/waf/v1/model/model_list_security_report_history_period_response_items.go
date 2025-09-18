package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSecurityReportHistoryPeriodResponseItems struct {

	// **参数解释：** 报告ID，唯一标识历史统计周期对应的安全报告。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ReportId string `json:"report_id"`

	// **参数解释：** 订阅ID，关联历史统计周期所属的安全报告订阅。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubscriptionId string `json:"subscription_id"`

	StatPeriod *ListSecurityReportHistoryPeriodResponseStatPeriod `json:"stat_period"`
}

func (o ListSecurityReportHistoryPeriodResponseItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityReportHistoryPeriodResponseItems struct{}"
	}

	return strings.Join([]string{"ListSecurityReportHistoryPeriodResponseItems", string(data)}, " ")
}
