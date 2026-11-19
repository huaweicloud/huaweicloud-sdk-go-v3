package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ArchivableTable struct {

	// **参数解释**：  数据库名。  **取值范围**：  不涉及。
	DbName string `json:"db_name"`

	// **参数解释**：  表名。  **取值范围**：  不涉及。
	TableName string `json:"table_name"`

	// **参数解释**：  分区名，为空表示无分区表。  **取值范围**：  不涉及。
	PartitionName *string `json:"partition_name,omitempty"`

	// **参数解释**：  是否存在限制或约束。  **取值范围**：  - true：存在限制或约束，归档受限。 - false：无限制或约束，可归档。
	Restricted bool `json:"restricted"`
}

func (o ArchivableTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArchivableTable struct{}"
	}

	return strings.Join([]string{"ArchivableTable", string(data)}, " ")
}
