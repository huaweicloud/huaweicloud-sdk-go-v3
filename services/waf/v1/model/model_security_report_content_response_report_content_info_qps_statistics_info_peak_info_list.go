package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityReportContentResponseReportContentInfoQpsStatisticsInfoPeakInfoList struct {

	// **参数解释：** 统计维度标识（如ACCESS表示访问类统计）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释：** 时间线数据，按时间顺序排列的峰值QPS值。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Timeline *[]SecurityReportContentResponseReportContentInfoQpsStatisticsInfoTimeline1 `json:"timeline,omitempty"`
}

func (o SecurityReportContentResponseReportContentInfoQpsStatisticsInfoPeakInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentInfoQpsStatisticsInfoPeakInfoList struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentInfoQpsStatisticsInfoPeakInfoList", string(data)}, " ")
}
