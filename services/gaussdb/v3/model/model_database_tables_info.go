package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseTablesInfo 查询的数据库以及对应表名称的结构体
type DatabaseTablesInfo struct {

	// **参数解释**：  查询的数据库名称。  **约束限制**：  支持英文大小写字母、数字、下划线。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Database *string `json:"database,omitempty"`

	// **参数解释**：  查询的数据表名称。  **约束限制**：  支持英文大小写字母、数字、下划线。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Tables *[]string `json:"tables,omitempty"`
}

func (o DatabaseTablesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseTablesInfo struct{}"
	}

	return strings.Join([]string{"DatabaseTablesInfo", string(data)}, " ")
}
