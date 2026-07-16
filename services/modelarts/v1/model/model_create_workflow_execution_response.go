package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowExecutionResponse Response Object
type CreateWorkflowExecutionResponse struct {

	// **参数解释**：创建时间。 **取值范围**：不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**：执行记录名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：工作流执行ID。[获取方法请参见[获取Execution列表](ListWorkflowExecutions.xml)。](tag:hc)。 **取值范围**：不涉及。
	ExecutionId *string `json:"execution_id,omitempty"`

	// **参数解释**：执行记录描述。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：执行记录状态。 **取值范围**：枚举值如下： - init：初始化 - running：运行中 - completed：运行成功 - stopped：已停止 - abnormal：异常
	Status *string `json:"status,omitempty"`

	// **参数解释**：工作空间ID。获取方法请参见[查询工作空间列表](ListWorkspace.xml)。 **取值范围**：不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：Workflow工作流ID。 **取值范围**：不涉及。
	WorkflowId *string `json:"workflow_id,omitempty"`

	// **参数解释**：工作流名称。 **取值范围**：不涉及。
	WorkflowName *string `json:"workflow_name,omitempty"`

	// **参数解释**：自定义场景ID，[获取方法请参见[查询工作流执行记录列表](CreateWorkflow.xml)](tag:hc)。 **取值范围**：不涉及。
	SceneId *string `json:"scene_id,omitempty"`

	// **参数解释**：自定义场景名称。 **取值范围**：不涉及。
	SceneName *string `json:"scene_name,omitempty"`

	// **参数解释**：执行记录的step。 **约束限制**：不涉及。
	StepsExecution *[]StepExecutionResp `json:"steps_execution,omitempty"`

	// **参数解释**：子图。
	SubGraphs *[]WorkflowSubgraphResp `json:"sub_graphs,omitempty"`

	// **参数解释**：执行的时长。 **取值范围**：不涉及。
	Duration *string `json:"duration,omitempty"`

	// **参数解释**：执行的事件。
	Events *[]string `json:"events,omitempty"`

	// **参数解释**：为执行记录设置的标签。
	Labels *[]string `json:"labels,omitempty"`

	// **参数解释**：节点steps使用到的数据。
	DataRequirements *[]DataRequirementResp `json:"data_requirements,omitempty"`

	// **参数解释**：节点steps使用到的参数。
	Parameters *[]WorkflowParameterResp `json:"parameters,omitempty"`

	Policies       *WorkflowDagPoliciesResp `json:"policies,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CreateWorkflowExecutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowExecutionResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkflowExecutionResponse", string(data)}, " ")
}
