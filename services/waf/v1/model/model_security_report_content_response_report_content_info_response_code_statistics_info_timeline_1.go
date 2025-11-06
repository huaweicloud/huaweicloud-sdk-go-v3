package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityReportContentResponseReportContentInfoResponseCodeStatisticsInfoTimeline1 struct {

	// **参数解释：** 时间戳（毫秒级），标识统计数据对应的时间点。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Time *int64 `json:"time,omitempty"`

	// **参数解释：** 该时间点上游服务器返回该响应码的数量。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Num *int64 `json:"num,omitempty"`
}

func (o SecurityReportContentResponseReportContentInfoResponseCodeStatisticsInfoTimeline1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentInfoResponseCodeStatisticsInfoTimeline1 struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentInfoResponseCodeStatisticsInfoTimeline1", string(data)}, " ")
}
