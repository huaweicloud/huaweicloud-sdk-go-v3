package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetainedBackupInfo 回收站保留备份信息。
type RetainedBackupInfo struct {

	// **参数解释**：  备份ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	BackupId *string `json:"backup_id,omitempty"`

	// **参数解释**：  源实例ID（已释放实例）。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  备份状态。  **约束限制**：  不涉及。  **取值范围**：  Active：可用。  **默认取值**：  不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：  备份名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	BackupName *string `json:"backup_name,omitempty"`

	// **参数解释**：  备份大小，单位MB。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释**：  备份开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**：  备份结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”，且大于备份开始时间。  其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**：  备份类型。  **约束限制**：  不涉及。  **取值范围**：  - auto：自动全量备份。 - manual：手动全量备份。 - differential：差量备份。 - sparse：稀疏备份。 - recycle: 回收站备份。  **默认取值**：  不涉及。
	BackupType *string `json:"backup_type,omitempty"`

	// **参数解释**：  备份级别，即备份位置。  **约束限制**：  不涉及。  **取值范围**：  - 1：一级备份。 - 2：二级备份。  **默认取值**：  不涉及。
	BackupLevel *string `json:"backup_level,omitempty"`
}

func (o RetainedBackupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetainedBackupInfo struct{}"
	}

	return strings.Join([]string{"RetainedBackupInfo", string(data)}, " ")
}
