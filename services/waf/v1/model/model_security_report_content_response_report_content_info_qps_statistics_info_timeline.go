package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityReportContentResponseReportContentInfoQpsStatisticsInfoTimeline struct {

	// **参数解释：** 时间戳（毫秒级），标识统计数据对应的时间点。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Time *int64 `json:"time,omitempty"`

	// **参数解释：** 该时间点对应统计维度的平均QPS值。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Num *int64 `json:"num,omitempty"`
}

func (o SecurityReportContentResponseReportContentInfoQpsStatisticsInfoTimeline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentInfoQpsStatisticsInfoTimeline struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentInfoQpsStatisticsInfoTimeline", string(data)}, " ")
}
