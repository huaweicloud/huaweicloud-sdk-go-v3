package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BotRequestTimeline **参数解释：** 特定时间点的请求趋势数据 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type BotRequestTimeline struct {

	// **参数解释：** 时间点（如2023-10-01 00:00:00） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Datetime *int64 `json:"datetime,omitempty"`

	// **参数解释：** 该时间点的正常请求数量 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	NormalRequestCount *int32 `json:"normal_request_count,omitempty"`

	// **参数解释：** 该时间点匹配已知bot的请求数量 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	KnownBotMatchedCount *int32 `json:"known_bot_matched_count,omitempty"`

	// **参数解释：** 该时间点透明检测的请求数量 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	TransparentMatchedCount *int32 `json:"transparent_matched_count,omitempty"`

	// **参数解释：** 该时间点行为检测的请求数量 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	BehaviorMatchedCount *int32 `json:"behavior_matched_count,omitempty"`
}

func (o BotRequestTimeline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BotRequestTimeline struct{}"
	}

	return strings.Join([]string{"BotRequestTimeline", string(data)}, " ")
}
