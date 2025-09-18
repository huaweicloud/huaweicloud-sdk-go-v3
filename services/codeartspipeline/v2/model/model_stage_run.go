package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StageRun 阶段运行信息
type StageRun struct {

	// **参数解释**： 阶段ID。 **取值范围**： 32位字符，由数字和字母组成。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 阶段类型。 **取值范围**： 不涉及。
	Category *string `json:"category,omitempty"`

	// **参数解释**： 阶段名称。 **取值范围**： 仅支持输入中文、大小写英文字母、数字、'-'、'_'、','、';'、':'、'.'、'/'、'('、')'、'（'、'）'及空格，其中空格不可在名称开头或结尾使用，且长度为[1,128]个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 阶段唯一标识。 **取值范围**： 不涉及。
	Identifier *string `json:"identifier,omitempty"`

	// **参数解释**： 是否总是运行。 **取值范围**： - true：总是运行。 - false：非总是运行。
	RunAlways *bool `json:"run_always,omitempty"`

	// **参数解释**： 是否并行。 **取值范围**： 不涉及。
	Parallel *string `json:"parallel,omitempty"`

	// **参数解释**： 是否选中。 **取值范围**： - true：选中。 - false：未选中。
	IsSelect *bool `json:"is_select,omitempty"`

	// **参数解释**： 序列号。 **取值范围**： 大于等于0。
	Sequence *int32 `json:"sequence,omitempty"`

	// **参数解释**： 依赖阶段的identifier信息。 **取值范围**： 不涉及。
	DependsOn *[]string `json:"depends_on,omitempty"`

	// **参数解释**： 运行条件。 **取值范围**： 不涉及。
	Condition *string `json:"condition,omitempty"`

	// **参数解释**： 状态。 **取值范围**： - INIT：初始化。 - QUEUED：排队。 - RUNNING：运行中。 - CANCELED：取消。 - COMPLETED：已完成。 - FAILED：失败。 - SKIPPED：跳过。 - IGNORED：忽略。 - PAUSED：暂停。 - SUSPEND：挂起。 - ASYNC_RUNNING：异步运行。 - ASYNC_FAILED：异步失败。 - UNSELECTED：未选择。 - REDISPATCH：重新调度。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 阶段开始时间。 **取值范围**： 不涉及。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 阶段结束时间。 **取值范围**： 不涉及。
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**： 阶段准入。 **取值范围**： 不涉及。
	Pre *[]StepRun `json:"pre,omitempty"`

	// **参数解释**： 阶段准出。 **取值范围**： 不涉及。
	Post *[]StepRun `json:"post,omitempty"`

	// **参数解释**： 任务列表。 **取值范围**： 不涉及。
	Jobs *[]JobRun `json:"jobs,omitempty"`
}

func (o StageRun) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StageRun struct{}"
	}

	return strings.Join([]string{"StageRun", string(data)}, " ")
}
