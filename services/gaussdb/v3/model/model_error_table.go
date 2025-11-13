package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErrorTable struct {

	// **参数解释**：  Excel导入失败的对象信息。  **取值范围**：  不涉及。
	Message *string `json:"message,omitempty"`

	// **参数解释**：  Excel导入失败的数据库名。   **取值范围**：  不涉及。
	DatabaseName *string `json:"database_name,omitempty"`

	// **参数解释**：  Excel导入失败的表名。   **取值范围**：  不涉及。
	TableName *string `json:"table_name,omitempty"`

	// **参数解释**：  Excel导入失败的行。  **取值范围**：  不涉及。
	RowNumber *int32 `json:"row_number,omitempty"`

	// **参数解释**：  Excel导入失败的错误信息描述。  **取值范围**：  不涉及。
	FullDescription *string `json:"full_description,omitempty"`
}

func (o ErrorTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorTable struct{}"
	}

	return strings.Join([]string{"ErrorTable", string(data)}, " ")
}
