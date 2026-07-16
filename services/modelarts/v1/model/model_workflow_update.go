package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowUpdate 工作流更新信息。
type WorkflowUpdate struct {

	// 工作流名称。填写1-64位，仅包含英文、数字、下划线（_）和中划线（-），并且以英文开头的名称。
	Name *string `json:"name,omitempty"`

	// 工作流描述。
	Description *string `json:"description,omitempty"`

	// Workflow包含的数据输入项定义。
	DataRequirements *[]DataRequirement `json:"data_requirements,omitempty"`

	// 工作流参数。
	Parameters *[]WorkflowParameter `json:"parameters,omitempty"`

	// 工作流存储信息。
	Storages *[]WorkflowStorage `json:"storages,omitempty"`

	// 工作流标签。
	Labels *[]string `json:"labels,omitempty"`

	// SMN消息订阅开关。
	SmnSwitch *string `json:"smn_switch,omitempty"`

	// 工作流节点。
	Steps *[]WorkflowStep `json:"steps,omitempty"`
}

func (o WorkflowUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowUpdate struct{}"
	}

	return strings.Join([]string{"WorkflowUpdate", string(data)}, " ")
}
