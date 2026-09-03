package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetainBackup **参数解释**  保留的备份  **约束限制**  不涉及  **取值范围**  不涉及  **默认取值**  不涉及
type RetainBackup struct {

	// **参数解释**：  备份名字  **约束限制**  不涉及  **取值范围**  不涉及  **默认取值**  不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：  备份ID  **约束限制**  不涉及  **取值范围**  不涉及  **默认取值**  不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**：  备份类型。Db表示自动备份、Snapshot表示手动备份  **约束限制**  不涉及  **取值范围**  Db、Snapshot  **默认取值**  不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**：  备份开始时间  **约束限制**  不涉及  **取值范围**  不涉及  **默认取值**  不涉及
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**：  备份结束时间  **约束限制**  不涉及  **取值范围**  不涉及  **默认取值**  不涉及
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**：  备份大小  **约束限制**  不涉及  **取值范围**  不涉及  **默认取值**  不涉及
	Size *string `json:"size,omitempty"`

	// **参数解释**：  备份描述信息  **约束限制**  不涉及  **取值范围**  不涉及  **默认取值**  不涉及
	Describe *string `json:"describe,omitempty"`

	// **参数解释**：  备份方式。Physics表示物理备份、Snapshot表示快照备份  **约束限制**  不涉及  **取值范围**  Physics、Snapshot  **默认取值**  不涉及
	BackupMethod *string `json:"backup_method,omitempty"`

	// **参数解释**：  备份是否tde加密  **约束限制**  不涉及  **取值范围**  false、true  **默认取值**  不涉及
	Tde *bool `json:"tde,omitempty"`
}

func (o RetainBackup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetainBackup struct{}"
	}

	return strings.Join([]string{"RetainBackup", string(data)}, " ")
}
