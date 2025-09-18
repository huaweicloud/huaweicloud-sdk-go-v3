package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfo **参数解释：** 带宽统计信息，包含平均带宽和峰值带宽的各维度时间线统计。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfo struct {

	// **参数解释：** 平均带宽统计列表，包含各维度按时间线的平均带宽数据。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AverageInfoList *[]SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfoAverageInfoList `json:"average_info_list,omitempty"`

	// **参数解释：** 峰值带宽统计列表，包含各维度按时间线的峰值带宽数据。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PeakInfoList *[]SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfoPeakInfoList `json:"peak_info_list,omitempty"`
}

func (o SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfo struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfo", string(data)}, " ")
}
