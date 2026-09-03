package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreTablesInfo 已恢复库表信息
type RestoreTablesInfo struct {

	// **参数解释**：  数据库名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	DatabaseName *string `json:"database_name,omitempty"`

	// **参数解释**：  表名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	TableName *string `json:"table_name,omitempty"`
}

func (o RestoreTablesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreTablesInfo struct{}"
	}

	return strings.Join([]string{"RestoreTablesInfo", string(data)}, " ")
}
