package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBotMTimelineResponse Response Object
type ListBotMTimelineResponse struct {

	// **参数解释：** 统计开始时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释：** 统计结束时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释：** 时间间隔（如1h表示每小时） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TimeSpan *int64 `json:"time_span,omitempty"`

	Timelines      *[]BotRequestTimeline `json:"timelines,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListBotMTimelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBotMTimelineResponse struct{}"
	}

	return strings.Join([]string{"ListBotMTimelineResponse", string(data)}, " ")
}
