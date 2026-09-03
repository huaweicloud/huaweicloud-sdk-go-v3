package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBackupUsageAlarmConfigRequestBody 修改备份空间告警配置请求体
type UpdateBackupUsageAlarmConfigRequestBody struct {

	// **参数解释**：  告警开关。  **约束限制**：  不涉及。  **取值范围**：  - ON - OFF  **默认取值**：  不涉及。
	AlarmEnabled string `json:"alarm_enabled"`

	// **参数解释**：  阈值百分比，占免费备份空间大小的百分比。  **约束限制**：  不涉及。  **取值范围**：  1-100。  **默认取值**：  90
	ThresholdPercent *int32 `json:"threshold_percent,omitempty"`

	// **参数解释**：  增量百分比，占免费备份空间大小的百分比。  **约束限制**：  不涉及。  **取值范围**：  1-100。  **默认取值**：  10
	IncrementPercent *int32 `json:"increment_percent,omitempty"`
}

func (o UpdateBackupUsageAlarmConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupUsageAlarmConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateBackupUsageAlarmConfigRequestBody", string(data)}, " ")
}
