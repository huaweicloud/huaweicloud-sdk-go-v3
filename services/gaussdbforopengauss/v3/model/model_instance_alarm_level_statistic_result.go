package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceAlarmLevelStatisticResult struct {

	// **参数解释**: 实例ID。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 实例名称。 **取值范围**: 不涉及。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**: 告警总数。 **取值范围**: 不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**: 实例级别的各等级告警数量统计。
	AlarmLevelStatistics *[]AlarmLevelStatisticsResult `json:"alarm_level_statistics,omitempty"`
}

func (o InstanceAlarmLevelStatisticResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceAlarmLevelStatisticResult struct{}"
	}

	return strings.Join([]string{"InstanceAlarmLevelStatisticResult", string(data)}, " ")
}
