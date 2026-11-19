package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ColdTableWaitingTask struct {

	// **参数解释**：  数据库名。  **取值范围**：  不涉及。
	DbName *string `json:"db_name,omitempty"`

	// **参数解释**：  表名。  **取值范围**：  不涉及。
	TableName *string `json:"table_name,omitempty"`

	// **参数解释**：  分区名。  **取值范围**：  不涉及。
	PartitionName *string `json:"partition_name,omitempty"`

	// **参数解释**：  任务状态。  **取值范围**：  - waiting_archive：等待归档。 - waiting_unarchive：等待回迁。
	Status *string `json:"status,omitempty"`

	// **参数解释**：  任务创建时间，13位毫秒级时间戳。  **取值范围**：  非负整数。
	CreateAt *int64 `json:"create_at,omitempty"`
}

func (o ColdTableWaitingTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColdTableWaitingTask struct{}"
	}

	return strings.Join([]string{"ColdTableWaitingTask", string(data)}, " ")
}
