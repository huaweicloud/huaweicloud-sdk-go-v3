package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkItemFlowProcessInstanceVo 工作流实例信息
type WorkItemFlowProcessInstanceVo struct {

	// **参数解释**： 工作项工作流实例ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 工作流实例是否挂起。 **取值范围**：  1: 运行  2: 挂起
	FlowState *int32 `json:"flow_state,omitempty"`

	// **参数解释**： 工作流入口ID。 **取值范围**： 不涉及。
	WorkflowEntryId *string `json:"workflow_entry_id,omitempty"`

	// **参数解释**： 工作流分类。 **取值范围**： 不涉及。
	Category *string `json:"category,omitempty"`
}

func (o WorkItemFlowProcessInstanceVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemFlowProcessInstanceVo struct{}"
	}

	return strings.Join([]string{"WorkItemFlowProcessInstanceVo", string(data)}, " ")
}
