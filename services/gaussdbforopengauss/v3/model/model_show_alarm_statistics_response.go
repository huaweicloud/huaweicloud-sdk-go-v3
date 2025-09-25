package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmStatisticsResponse Response Object
type ShowAlarmStatisticsResponse struct {

	// **参数解释**: 告警总数。 **取值范围**: 不涉及。
	TotalAlarmCount *int32 `json:"total_alarm_count,omitempty"`

	// **参数解释**: 环比比率。 **取值范围**: 值为0表示环比没有变化，值为空表示上一周期没有告警。
	RingPercentage *float64 `json:"ring_percentage,omitempty"`

	// **参数解释**: 实例级别的告警统计。
	InstanceAlarmLevelStatistics *[]InstanceAlarmLevelStatisticResult `json:"instance_alarm_level_statistics,omitempty"`

	// **参数解释**: 全量告警统计。
	TotalAlarmLevelStatistics *[]AlarmLevelStatisticsResult `json:"total_alarm_level_statistics,omitempty"`
	HttpStatusCode            int                           `json:"-"`
}

func (o ShowAlarmStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowAlarmStatisticsResponse", string(data)}, " ")
}
