package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMigrationLogResponse Response Object
type ShowMigrationLogResponse struct {

	// **参数解释**：  分片变更任务日志的集合。  **参数范围**：  不涉及。
	TaskLogs *[]TaskLogsVo `json:"task_logs,omitempty"`

	// **参数解释**：  当前页。  **参数范围**：  不涉及。
	CurPage *int32 `json:"cur_page,omitempty"`

	// **参数解释**：  每页条数。  **参数范围**：  不涉及。
	PerPage *int32 `json:"per_page,omitempty"`

	// **参数解释**：  总量。  **参数范围**：  不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowMigrationLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrationLogResponse struct{}"
	}

	return strings.Join([]string{"ShowMigrationLogResponse", string(data)}, " ")
}
