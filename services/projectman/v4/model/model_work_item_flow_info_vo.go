package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkItemFlowInfoVo 工作项流转信息
type WorkItemFlowInfoVo struct {
	ProcessInstance *WorkItemFlowProcessInstanceVo `json:"process_instance,omitempty"`

	// **参数解释**： 工作项关联的全部工作流节点列表。 **取值范围**： 不涉及。
	ProcessNodes *[]WorkItemFlowProcessNodeVo `json:"process_nodes,omitempty"`

	CurrentProcessNode *WorkItemFlowProcessNodeVo `json:"current_process_node,omitempty"`

	// **参数解释**： 可以流转的流转线信息。 **取值范围**： 不涉及。
	NextFlow *[]FlowsInfoVo `json:"next_flow,omitempty"`

	// **参数解释**： 流转失败时的失败原因。 **取值范围**： 不涉及。
	FailResult *string `json:"fail_result,omitempty"`
}

func (o WorkItemFlowInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemFlowInfoVo struct{}"
	}

	return strings.Join([]string{"WorkItemFlowInfoVo", string(data)}, " ")
}
