package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfoPeakInfoList struct {

	// **参数解释：** 统计维度标识（如BANDWIDTH表示带宽类统计）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释：** 时间线数据，按时间顺序排列的峰值带宽值。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Timeline *[]SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfoTimeline1 `json:"timeline,omitempty"`
}

func (o SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfoPeakInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfoPeakInfoList struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfoPeakInfoList", string(data)}, " ")
}
