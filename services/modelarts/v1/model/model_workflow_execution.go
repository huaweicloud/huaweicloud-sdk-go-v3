package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkflowExecution struct {

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 执行记录名称。
	Name *string `json:"name,omitempty"`

	// 工作流执行ID。
	ExecutionId *string `json:"execution_id,omitempty"`

	// 执行记录描述。
	Description *string `json:"description,omitempty"`

	// 执行记录状态。
	Status *string `json:"status,omitempty"`

	// 工作空间ID。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// Workflow工作流ID。
	WorkflowId *string `json:"workflow_id,omitempty"`

	// 工作流名称。填写1-64位，仅包含英文、数字、下划线（_）和中划线（-），并且以英文开头的名称。
	WorkflowName *string `json:"workflow_name,omitempty"`

	// 自定义场景ID。
	SceneId *string `json:"scene_id,omitempty"`

	// 自定义场景名称。
	SceneName *string `json:"scene_name,omitempty"`

	// 执行记录的step。
	StepsExecution *[]StepExecution `json:"steps_execution,omitempty"`

	// 子图。
	SubGraphs *[]WorkflowSubgraph `json:"sub_graphs,omitempty"`

	// 执行的时长。
	Duration *string `json:"duration,omitempty"`

	// 执行的事件。
	Events *[]string `json:"events,omitempty"`

	// 为执行记录设置的标签。
	Labels *[]string `json:"labels,omitempty"`

	// 节点steps使用到的数据。
	DataRequirements *[]DataRequirement `json:"data_requirements,omitempty"`

	// 节点steps使用到的参数。
	Parameters *[]WorkflowParameter `json:"parameters,omitempty"`

	Policies *WorkflowDagPolicies `json:"policies,omitempty"`
}

func (o WorkflowExecution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowExecution struct{}"
	}

	return strings.Join([]string{"WorkflowExecution", string(data)}, " ")
}
