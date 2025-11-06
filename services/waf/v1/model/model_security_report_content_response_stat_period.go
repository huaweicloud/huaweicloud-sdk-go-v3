package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityReportContentResponseStatPeriod **参数解释：** 统计周期，标识当前安全报告统计数据的时间范围。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type SecurityReportContentResponseStatPeriod struct {

	// **参数解释：** 开始时间，统计周期的起始时间戳（毫秒级）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	BeginTime int64 `json:"begin_time"`

	// **参数解释：** 结束时间，统计周期的终止时间戳（毫秒级）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EndTime int64 `json:"end_time"`
}

func (o SecurityReportContentResponseStatPeriod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseStatPeriod struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseStatPeriod", string(data)}, " ")
}
