package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBackupUsageAlarmConfigResponse Response Object
type UpdateBackupUsageAlarmConfigResponse struct {

	// **参数解释**：  操作执行状态。  **约束限制**：  不涉及。  **取值范围**：  - COMPLETED  **默认取值**：  不涉及。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateBackupUsageAlarmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupUsageAlarmConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateBackupUsageAlarmConfigResponse", string(data)}, " ")
}
