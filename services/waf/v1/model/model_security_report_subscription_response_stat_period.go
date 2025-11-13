package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityReportSubscriptionResponseStatPeriod **参数解释：** 统计周期，订阅报告的默认统计时间范围配置。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type SecurityReportSubscriptionResponseStatPeriod struct {

	// **参数解释：** 开始时间，订阅报告统计周期的起始时间戳（毫秒级）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	BeginTime *int64 `json:"begin_time,omitempty"`

	// **参数解释：** 结束时间，订阅报告统计周期的终止时间戳（毫秒级）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o SecurityReportSubscriptionResponseStatPeriod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportSubscriptionResponseStatPeriod struct{}"
	}

	return strings.Join([]string{"SecurityReportSubscriptionResponseStatPeriod", string(data)}, " ")
}
