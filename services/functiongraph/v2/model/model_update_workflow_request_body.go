package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 函数流创建body体
type UpdateWorkflowRequestBody struct {
	// 流程定义名称

	Name *string `json:"name,omitempty"`
	// 流程定义描述

	Description *string `json:"description,omitempty"`
	// 触发器列表

	Triggers *[]Trigger `json:"triggers,omitempty"`
	// 流程开始节点ID

	Start *string `json:"start,omitempty"`
	// 函数清单

	Functions *[]Function `json:"functions,omitempty"`
	// 工作流节点清单，定义参考SleepState和OperationState

	States *[]OperationState `json:"states,omitempty"`
	// 工作流中的常量

	Constants *interface{} `json:"constants,omitempty"`
	// 重试策略清单

	Retries *[]Retry `json:"retries,omitempty"`
	// 企业项目ID，在企业用户创建函数时必填。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o UpdateWorkflowRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkflowRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateWorkflowRequestBody", string(data)}, " ")
}
