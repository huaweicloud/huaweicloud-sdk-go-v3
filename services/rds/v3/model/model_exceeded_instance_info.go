package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExceededInstanceInfo 超阈值实例信息
type ExceededInstanceInfo struct {

	// **参数解释**：  实例ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  日志备份空间使用量，单位GB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ObsUsageGb *float64 `json:"obs_usage_gb,omitempty"`

	// **参数解释**：  日志备份免费备份空间额度，单位GB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ObsFreeBackupSpaceGb *float64 `json:"obs_free_backup_space_gb,omitempty"`

	// **参数解释**：  快照备份空间使用量，单位GB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SnapshotUsageGb *float64 `json:"snapshot_usage_gb,omitempty"`

	// **参数解释**：  快照免费备份空间额度，单位GB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SnapshotFreeBackupSpaceGb *float64 `json:"snapshot_free_backup_space_gb,omitempty"`
}

func (o ExceededInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExceededInstanceInfo struct{}"
	}

	return strings.Join([]string{"ExceededInstanceInfo", string(data)}, " ")
}
