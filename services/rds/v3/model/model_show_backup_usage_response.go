package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupUsageResponse Response Object
type ShowBackupUsageResponse struct {

	// **参数解释**：  备份总使用量，各类备份占用的备份总大小。 单位：MB  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	BackupUseSpace *float64 `json:"backup_use_space,omitempty"`

	// **参数解释**：  物理备份总使用量，包括本区域的物理全量备份，binlog日志备份，审计日志。 单位：MB  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	DbUseSpace *float64 `json:"db_use_space,omitempty"`

	// **参数解释**：  由RDS计费的cbr快照备份总使用量， 单位MB  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	RdsSnapshotUseSpace *float64 `json:"rds_snapshot_use_space,omitempty"`

	// **参数解释**：  跨区域备份总使用量，包括跨区域的物理全量备份，binlog日志备份， 单位：MB  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	OffsiteUseSpace *float64 `json:"offsite_use_space,omitempty"`
	HttpStatusCode  int      `json:"-"`
}

func (o ShowBackupUsageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupUsageResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupUsageResponse", string(data)}, " ")
}
