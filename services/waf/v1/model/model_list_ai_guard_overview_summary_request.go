package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiGuardOverviewSummaryRequest Request Object
type ListAiGuardOverviewSummaryRequest struct {

	// **参数解释：** 开始时间戳（毫秒），必填。时间范围不超过1个月 **约束限制：** 必填 **取值范围：** 不涉及 **默认取值：** 不涉及
	From int64 `json:"from"`

	// **参数解释：** 结束时间戳（毫秒），必填。时间范围不超过1个月 **约束限制：** 必填 **取值范围：** 不涉及 **默认取值：** 不涉及
	To int64 `json:"to"`
}

func (o ListAiGuardOverviewSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiGuardOverviewSummaryRequest struct{}"
	}

	return strings.Join([]string{"ListAiGuardOverviewSummaryRequest", string(data)}, " ")
}
