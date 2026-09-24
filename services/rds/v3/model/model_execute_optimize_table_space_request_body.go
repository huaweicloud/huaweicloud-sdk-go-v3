package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteOptimizeTableSpaceRequestBody 清理表碎片空间的请求体。
type ExecuteOptimizeTableSpaceRequestBody struct {

	// **参数解释**：  表名。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	TableName string `json:"table_name"`

	// **参数解释**：  数据库名。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	DatabaseName string `json:"database_name"`
}

func (o ExecuteOptimizeTableSpaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteOptimizeTableSpaceRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteOptimizeTableSpaceRequestBody", string(data)}, " ")
}
