package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CancelColdTableWaitingTask struct {

	// **参数解释**：  数据库名。  获取方法请参见[查询冷热分离待下发任务列表](https://support.huaweicloud.com/api-taurusdb/ListColdTableWaitingTasks.html)。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	DbName string `json:"db_name"`

	// **参数解释**：  表名。 获取方法请参见[查询冷热分离待下发任务列表](https://support.huaweicloud.com/api-taurusdb/ListColdTableWaitingTasks.html)。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	TableName string `json:"table_name"`

	// **参数解释**：  分区名。  获取方法请参见[查询冷热分离待下发任务列表](https://support.huaweicloud.com/api-taurusdb/ListColdTableWaitingTasks.html)。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	PartitionName *string `json:"partition_name,omitempty"`

	// **参数解释**：  任务类型。  **约束限制**：  不涉及。  **取值范围**：  - waiting_archive：待归档。 - waiting_unarchive：待回迁。  **默认取值**：  不涉及。
	Type string `json:"type"`
}

func (o CancelColdTableWaitingTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelColdTableWaitingTask struct{}"
	}

	return strings.Join([]string{"CancelColdTableWaitingTask", string(data)}, " ")
}
