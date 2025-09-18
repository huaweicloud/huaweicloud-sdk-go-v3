package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobRun 任务运行信息
type JobRun struct {

	// **参数解释**： 任务ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 任务类型。 **取值范围**： 不涉及。
	Category *string `json:"category,omitempty"`

	// **参数解释**： 序列号。 **取值范围**： 不涉及。
	Sequence *int32 `json:"sequence,omitempty"`

	// **参数解释**： 是否异步。 **取值范围**： - true：异步。 - false：非异步。
	Async *string `json:"async,omitempty"`

	// **参数解释**： 任务名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 任务唯一标识。 **取值范围**： 不涉及。
	Identifier *string `json:"identifier,omitempty"`

	// **参数解释**： 依赖。 **取值范围**： 不涉及。
	DependsOn *[]string `json:"depends_on,omitempty"`

	// **参数解释**： 运行条件。 **取值范围**： 不涉及。
	Condition *string `json:"condition,omitempty"`

	// **参数解释**： 执行资源。 **取值范围**： 不涉及。
	Resource *string `json:"resource,omitempty"`

	// **参数解释**： 是否选中。 **取值范围**： - true：选中。 - false：未选中。
	IsSelect *bool `json:"is_select,omitempty"`

	// **参数解释**： 任务超时设置。 **取值范围**： 不涉及。
	Timeout *string `json:"timeout,omitempty"`

	// **参数解释**： 任务上次下发ID。 **取值范围**： 不涉及。
	LastDispatchId *string `json:"last_dispatch_id,omitempty"`

	// **参数解释**： 状态。 **取值范围**： - INIT：初始化。 - QUEUED：排队。 - RUNNING：运行中。 - CANCELED：取消。 - COMPLETED：已完成。 - FAILED：失败。 - SKIPPED：跳过。 - IGNORED：忽略。 - PAUSED：暂停。 - SUSPEND：挂起。 - ASYNC_RUNNING：异步运行。 - ASYNC_FAILED：异步失败。 - UNSELECTED：未选择。 - REDISPATCH：重新调度。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 错误信息。 **取值范围**： 不涉及。
	Message *string `json:"message,omitempty"`

	// **参数解释**： 任务开始时间。 **取值范围**： 不涉及。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 任务结束时间。 **取值范围**： 不涉及。
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： 步骤。 **取值范围**： 不涉及。
	Steps *[]StepRun `json:"steps,omitempty"`

	// **参数解释**： 任务执行ID。 **取值范围**： 不涉及。
	ExecId *string `json:"exec_id,omitempty"`
}

func (o JobRun) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobRun struct{}"
	}

	return strings.Join([]string{"JobRun", string(data)}, " ")
}
