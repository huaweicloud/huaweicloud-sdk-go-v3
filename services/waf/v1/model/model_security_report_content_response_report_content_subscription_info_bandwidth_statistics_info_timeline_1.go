package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfoTimeline1 struct {

	// **参数解释：** 时间戳（毫秒级），标识统计数据对应的时间点。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Time *int64 `json:"time,omitempty"`

	// **参数解释：** 该时间点对应统计维度的峰值带宽值。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Num *int32 `json:"num,omitempty"`
}

func (o SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfoTimeline1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfoTimeline1 struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentSubscriptionInfoBandwidthStatisticsInfoTimeline1", string(data)}, " ")
}
