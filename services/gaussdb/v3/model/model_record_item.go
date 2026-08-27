package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecordItem **参数解释**：  单个无锁变更任务信息。  **取值范围**：  不涉及。
type RecordItem struct {

	// **参数解释**：  无锁变更任务记录标识。  **取值范围**：  不涉及。
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**：  无锁变更任务详细内容。
	TaskContent *[]OnlineDdlTaskContentItem `json:"task_content,omitempty"`

	// **参数解释**：  无锁变更任务创建时间，13位毫秒时间戳。  **取值范围**： 不涉及。
	CreatedAt *int64 `json:"created_at,omitempty"`

	// **参数解释**：  无锁变更任务结束时间，13位毫秒时间戳。  **取值范围**： 不涉及。
	EndedAt *int64 `json:"ended_at,omitempty"`

	// **参数解释**：  无锁变更任务执行状态。  **取值范围**：   - checking：表示正在执行预检查步骤。   - check successful： 表示预检查步骤执行成功。   - check failed： 表示预检查步骤执行失败。   - altering： 表示正在任务正在执行变更步骤。   - alter successful： 表示变更步骤执行成功。   - alter failed： 表示变更步骤执行失败。   - stopping：表示正在执行停止任务步骤。   - stop successful： 表示执行停止步骤成功。   - stop failed： 表示执行停止步骤失败。   - cleaning： 表示正在执行清理临时表步骤。   - clean successful： 表示清理临时表步骤执行成功。   - clean failed： 表示清理临时表步骤执行失败。
	TaskStatus *string `json:"task_status,omitempty"`

	// **参数解释**：  表示数据库内核层面无锁变更任务运行阶段。  **取值范围**：   - 0：表示无锁变更任务未开始。  - 1：表示无锁变更任务已完成资源初始化。  - 2：表示无锁变更任务正在运行。  - 3：表示无锁变更任务已完成。
	AlterStage *int32 `json:"alter_stage,omitempty"`

	// **参数解释**：  无锁变更任务百分比进度，1位小数精度。  **取值范围**：  0.0-100.0。
	Percentage *float32 `json:"percentage,omitempty"`

	// **参数解释**：  无锁变更任务失败原因，任务执行失败时有返回值。  **取值范围**： 不涉及。
	ErrorReason *string `json:"error_reason,omitempty"`

	// **参数解释**：  无锁变更任务临时表名称，关闭临时表自动清理时有返回值。  **取值范围**： 不涉及。
	TempTableName *string `json:"temp_table_name,omitempty"`
}

func (o RecordItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordItem struct{}"
	}

	return strings.Join([]string{"RecordItem", string(data)}, " ")
}
