package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatisticPeriod **参数解释**： 统计周期，自定义报告需要设置 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type StatisticPeriod struct {

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： 结束时间 **约束限制**： 不涉及 **取值范围**： 毫秒级时间戳 **默认取值**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`
}

func (o StatisticPeriod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticPeriod struct{}"
	}

	return strings.Join([]string{"StatisticPeriod", string(data)}, " ")
}
