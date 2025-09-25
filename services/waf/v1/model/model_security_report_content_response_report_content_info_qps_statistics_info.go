package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityReportContentResponseReportContentInfoQpsStatisticsInfo **参数解释：** QPS统计信息，包含平均QPS和峰值QPS的各维度时间线统计。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type SecurityReportContentResponseReportContentInfoQpsStatisticsInfo struct {

	// **参数解释：** 平均QPS统计列表，包含各维度按时间线的平均QPS数据。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AverageInfoList *[]SecurityReportContentResponseReportContentInfoQpsStatisticsInfoAverageInfoList `json:"average_info_list,omitempty"`

	// **参数解释：** 峰值QPS统计列表，包含各维度按时间线的峰值QPS数据。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PeakInfoList *[]SecurityReportContentResponseReportContentInfoQpsStatisticsInfoPeakInfoList `json:"peak_info_list,omitempty"`
}

func (o SecurityReportContentResponseReportContentInfoQpsStatisticsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentInfoQpsStatisticsInfo struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentInfoQpsStatisticsInfo", string(data)}, " ")
}
