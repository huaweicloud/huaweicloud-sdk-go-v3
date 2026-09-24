package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Space 备份用量详情
type Space struct {

	// **参数解释**：  日志备份用量，单位MB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Obs *float64 `json:"obs,omitempty"`

	// **参数解释**：  审计日志用量，单位MB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Auditlog *float64 `json:"auditlog,omitempty"`

	// **参数解释**：  rds侧快照备份用量，单位MB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Snapshot *float64 `json:"snapshot,omitempty"`

	// **参数解释**：  rds侧CBR快照备份用量，单位MB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	CbrSnapshot *float64 `json:"cbr_snapshot,omitempty"`

	// **参数解释**：  日志备份赠送空间，单位GB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ObsFree *float64 `json:"obs_free,omitempty"`

	// **参数解释**：  快照备份赠送空间，单位GB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SnapshotFree *float64 `json:"snapshot_free,omitempty"`

	// **参数解释**：  全量备份大小，单位MB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Db *float64 `json:"db,omitempty"`

	// **参数解释**：  增量备份大小，单位MB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Log *float64 `json:"log,omitempty"`
}

func (o Space) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Space struct{}"
	}

	return strings.Join([]string{"Space", string(data)}, " ")
}
