package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmLevelStatisticsResult struct {

	// **参数解释**: 告警数量。 **取值范围**: 不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**: 告警级别名称。 **取值范围**: - critical - major - minor - notice
	LevelName *string `json:"level_name,omitempty"`
}

func (o AlarmLevelStatisticsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmLevelStatisticsResult struct{}"
	}

	return strings.Join([]string{"AlarmLevelStatisticsResult", string(data)}, " ")
}
