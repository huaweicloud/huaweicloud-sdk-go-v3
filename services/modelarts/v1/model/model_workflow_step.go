package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowStep 工作流节点。
type WorkflowStep struct {

	// Workflow工作流节点的名称，在一个DAG中唯一，1到64位只包含中英文，数字，空格，下划线（_）和中划线（-），并且以中英文开头。
	Name string `json:"name"`

	// 节点的类型，枚举值如下: - job 训练 - labeling 标注 - release_dataset 数据集发布 - model 模型发布 - service 服务部署 - mrs_job MRS作业 - dataset_import 数据集导入 - create_dataset 创建数据集
	Type *string `json:"type,omitempty"`

	// 节点的输入项。
	Inputs *[]JobInput `json:"inputs,omitempty"`

	// 节点的输出项。
	Outputs *[]JobOutput `json:"outputs,omitempty"`

	// 节点的创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 工作流节点标题。
	Title *string `json:"title,omitempty"`

	// 节点的描述信息。
	Description *string `json:"description,omitempty"`

	// 节点属性。
	Properties map[string]interface{} `json:"properties,omitempty"`

	// 运行依赖的前置节点。
	DependSteps *[]string `json:"depend_steps,omitempty"`

	// 节点执行条件。
	Conditions *[]StepCondition `json:"conditions,omitempty"`

	// 条件节点分支。
	IfThenSteps *[]string `json:"if_then_steps,omitempty"`

	// 条件节点另一分支。
	ElseThenSteps *[]string `json:"else_then_steps,omitempty"`

	Policy *WorkflowStepPolicy `json:"policy,omitempty"`
}

func (o WorkflowStep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowStep struct{}"
	}

	return strings.Join([]string{"WorkflowStep", string(data)}, " ")
}
