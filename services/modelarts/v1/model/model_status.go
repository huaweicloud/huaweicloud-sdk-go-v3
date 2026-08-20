package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Status 训练作业状态信息。创建作业无需填写。
type Status struct {

	// 训练作业一级状态。可选值如下： - Creating：创建中 - Pending：等待中 - Running：运行中 - Failed：运行失败 - Completed：已完成 - Terminating：停止中 - Terminated：已停止 - Abnormal：异常
	Phase *string `json:"phase,omitempty"`

	// 训练作业二级状态为内部详细状态，可能会增加、修改、删除，不建议依赖。可选值如下： - Creating：创建中 - Queuing：排队中 - Running：运行中 - Failed：运行失败 - Completed：已完成 - Terminating：停止中 - Terminated：已停止 - CreateFailed：创建失败 - TerminatedFailed：停止失败 - Unknown：未知状态 - Lost：异常
	SecondaryPhase *string `json:"secondary_phase,omitempty"`

	// 训练作业运行时长，单位为毫秒。
	Duration *int64 `json:"duration,omitempty"`

	// 训练作业运行时节点数变化指标。
	NodeCountMetrics *[][]int32 `json:"node_count_metrics,omitempty"`

	// 训练作业子任务名称。
	Tasks *[]string `json:"tasks,omitempty"`

	// 训练作业开始时间，格式为时间戳。
	StartTime *int64 `json:"start_time,omitempty"`

	// 训练在子任务状态信息。
	TaskStatuses *[]TaskStatuses `json:"task_statuses,omitempty"`

	// 训练作业运行及故障恢复记录。
	RunningRecords *[]RunningRecord `json:"running_records,omitempty"`

	// **参数解释**：作业已经保留时长。  **约束限制**：仅当创建训练作业时，设置了`reserved_time`时返回。  **取值范围**：不涉及。    **默认取值**：不涉及。
	RetentionTime *int32 `json:"retention_time,omitempty"`

	// **参数解释**：训练作业各 Task 的 IP 信息。 **约束限制**：仅当查询请求携带 `host_ips` 时返回；且仅返回与筛选 IP 匹配的记录。 **取值范围**：不涉及。 **默认取值**：不传 `host_ips` 时不返回。
	TaskIps *[]TaskIp `json:"task_ips,omitempty"`
}

func (o Status) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Status struct{}"
	}

	return strings.Join([]string{"Status", string(data)}, " ")
}
