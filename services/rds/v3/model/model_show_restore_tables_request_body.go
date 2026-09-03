package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRestoreTablesRequestBody 获取已恢复库表信息请求体
type ShowRestoreTablesRequestBody struct {

	// **参数解释**：  任务流ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	WorkflowId string `json:"workflow_id"`

	// **参数解释**：  任务名称。  **约束限制**：  不涉及。  **取值范围**：  - RestoreTableMysql - RestoreDatabaseMysql  **默认取值**：  不涉及。
	WorkflowName string `json:"workflow_name"`
}

func (o ShowRestoreTablesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRestoreTablesRequestBody struct{}"
	}

	return strings.Join([]string{"ShowRestoreTablesRequestBody", string(data)}, " ")
}
