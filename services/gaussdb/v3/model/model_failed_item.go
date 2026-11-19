package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FailedItem struct {

	// **参数解释**：  数据库名。  **取值范围**：  不涉及。
	DbName string `json:"db_name"`

	// **参数解释**：  表名。  **取值范围**：  不涉及。
	TableName string `json:"table_name"`

	// **参数解释**：  分区名。  **取值范围**：  不涉及。
	PartitionName *string `json:"partition_name,omitempty"`

	// **参数解释**：  失败原因。  **取值范围**：  - restricted：内核限制无法下发。 - incorrect：库名、表名、分区名有误。 - duplicate：该表正在进行其他操作（待归档、待回迁、归档中、回迁中、已归档、已回迁）。
	Reason string `json:"reason"`
}

func (o FailedItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailedItem struct{}"
	}

	return strings.Join([]string{"FailedItem", string(data)}, " ")
}
