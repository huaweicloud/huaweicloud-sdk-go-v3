package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StepRun 步骤运行信息
type StepRun struct {

	// **参数解释**： 步骤名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 步骤插件名。 **取值范围**： 不涉及。
	Task *string `json:"task,omitempty"`

	// **参数解释**： 插件业务类型。 **取值范围**： - Normal：通用。 - Build：构建。 - Test：测试。 - Check：代码检查。 - Deploy：部署。
	BusinessType *string `json:"business_type,omitempty"`

	// **参数解释**： 输入参数。 **取值范围**： 不涉及。
	Inputs *[]StepRunInputs `json:"inputs,omitempty"`

	// **参数解释**： 序列号。 **取值范围**： 不涉及。
	Sequence *int32 `json:"sequence,omitempty"`

	// **参数解释**： 官方插件版本号。 **取值范围**： 不涉及。
	OfficialTaskVersion *string `json:"official_task_version,omitempty"`

	// **参数解释**： 唯一标识符。 **取值范围**： 不涉及。
	Identifier *string `json:"identifier,omitempty"`

	// **参数解释**： 是否可编辑。 **取值范围**： - true：可编辑。 - false：不可编辑。
	MultiStepEditable *int32 `json:"multi_step_editable,omitempty"`

	// **参数解释**： 步骤ID。 **取值范围**： 32位字符，由数字和字母组成。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 扩展点ID列表。 **取值范围**： 不涉及。
	EndpointIds *[]string `json:"endpoint_ids,omitempty"`

	// **参数解释**： 上次下发任务ID。 **取值范围**： 32位字符，由数字和字母组成。
	LastDispatchId *string `json:"last_dispatch_id,omitempty"`

	// **参数解释**： 状态。 **取值范围**： - RUNNING：运行中。 - CANCELED：取消。 - COMPLETED：已完成。 - FAILED：失败。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 错误消息。 **取值范围**： 不涉及。
	Message *string `json:"message,omitempty"`

	// **参数解释**： 步骤开始时间。 **取值范围**： 不涉及。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 步骤结束时间。 **取值范围**： 不涉及。
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o StepRun) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StepRun struct{}"
	}

	return strings.Join([]string{"StepRun", string(data)}, " ")
}
