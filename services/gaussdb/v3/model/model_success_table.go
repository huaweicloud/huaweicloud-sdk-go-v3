package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SuccessTable struct {

	// **参数解释**：  Excel导入成功的数据库名。   **取值范围**：  不涉及。
	DbName *string `json:"db_name,omitempty"`

	// **参数解释**：  Excel导入成功的表名。   **取值范围**：  不涉及。
	TableName *string `json:"table_name,omitempty"`

	// **参数解释**：  Excel导入成功的表配置。   **取值范围**：  不涉及。
	ConfigName *string `json:"config_name,omitempty"`

	// **参数解释**：  Excel导入成功的行数。   **取值范围**：  不涉及。
	RowNumber *int32 `json:"row_number,omitempty"`

	// **参数解释**：  Excel导入成功的表全名。   **取值范围**：  不涉及。
	FullTableName *string `json:"full_table_name,omitempty"`

	// **参数解释**：  Excel信息是否合规。   **取值范围**：  - true：校验结果合规。 - false：校验结果不合规。
	Valid *bool `json:"valid,omitempty"`
}

func (o SuccessTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SuccessTable struct{}"
	}

	return strings.Join([]string{"SuccessTable", string(data)}, " ")
}
