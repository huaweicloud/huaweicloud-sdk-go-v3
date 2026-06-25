package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceBackupSummary 实例备份概览
type InstanceBackupSummary struct {

	// **参数解释**：  实例ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  实例名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：  备份用量，单位MB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	BackupUseSpace *float64 `json:"backup_use_space,omitempty"`

	Datastore *InstanceBackupDatastore `json:"datastore,omitempty"`

	Spaces *Spaces `json:"spaces,omitempty"`
}

func (o InstanceBackupSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceBackupSummary struct{}"
	}

	return strings.Join([]string{"InstanceBackupSummary", string(data)}, " ")
}
