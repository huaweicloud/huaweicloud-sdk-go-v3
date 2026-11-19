package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ColdTableTask struct {

	// **参数解释**：  数据库名。  **取值范围**：  不涉及。
	DbName string `json:"db_name"`

	// **参数解释**：  表名。  **取值范围**：  不涉及。
	TableName string `json:"table_name"`

	// **参数解释**：  分区名，为空表示无分区表。  **取值范围**：  不涉及。
	PartitionName *string `json:"partition_name,omitempty"`

	// **参数解释**：  任务状态。  **取值范围**：  - BEGIN：开始。 - RUNNING：进行中。 - FINISH：已完成。 - FAIL：失败。 - REVOKE：回滚。
	Status string `json:"status"`

	// **参数解释**：  操作类型。  **取值范围**：  - OBS：创建冷表。 - SLICE：回迁冷表。
	TargetStorage *string `json:"target_storage,omitempty"`

	// **参数解释**：  任务进度百分比。  **取值范围**：  0-1。
	Percent *float64 `json:"percent,omitempty"`

	// **参数解释**：  对象数。  **取值范围**：  非负数。
	TotalCnt *float64 `json:"total_cnt,omitempty"`

	// **参数解释**：  表空间ID。  **取值范围**：  不涉及。
	SpaceId *string `json:"space_id,omitempty"`

	// **参数解释**：  冷表版本。  **取值范围**：  不涉及。
	Version *string `json:"version,omitempty"`
}

func (o ColdTableTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColdTableTask struct{}"
	}

	return strings.Join([]string{"ColdTableTask", string(data)}, " ")
}
