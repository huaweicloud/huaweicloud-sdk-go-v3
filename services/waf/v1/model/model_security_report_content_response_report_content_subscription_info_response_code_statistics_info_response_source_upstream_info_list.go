package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityReportContentResponseReportContentSubscriptionInfoResponseCodeStatisticsInfoResponseSourceUpstreamInfoList struct {

	// **参数解释：** 响应码标识（如504表示网关超时）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释：** 时间线数据，按时间顺序排列的上游响应码数量。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Timeline *[]SecurityReportContentResponseReportContentSubscriptionInfoResponseCodeStatisticsInfoTimeline1 `json:"timeline,omitempty"`
}

func (o SecurityReportContentResponseReportContentSubscriptionInfoResponseCodeStatisticsInfoResponseSourceUpstreamInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentSubscriptionInfoResponseCodeStatisticsInfoResponseSourceUpstreamInfoList struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentSubscriptionInfoResponseCodeStatisticsInfoResponseSourceUpstreamInfoList", string(data)}, " ")
}
