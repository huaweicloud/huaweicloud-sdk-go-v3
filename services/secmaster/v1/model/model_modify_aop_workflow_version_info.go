package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyAopWorkflowVersionInfo 修改流程版本的请求对象
type ModifyAopWorkflowVersionInfo struct {

	// **参数解释**: 流程名称 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**: 流程的描述 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**: 更新流程的动作 - pending_submit 更新流程基础信息 - pending_approval 提交审核 - not_activated 取消激活 - activated 激活流程  **约束限制**: 不涉及         **取值范围**: - pending_submit - pending_approval - not_activated - activated  **默认值**:  pending_submit
	Status *string `json:"status,omitempty"`

	// **参数解释**: 流程拓扑图的参数配置 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Taskconfig *string `json:"taskconfig,omitempty"`

	// **参数解释**: 流程的拓扑图的BASE64编码 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Taskflow *string `json:"taskflow,omitempty"`

	// **参数解释**: 流程的类型 **约束限制**: 不涉及         **取值范围**: - JSON  **默认值**:  不涉及
	TaskflowType *string `json:"taskflow_type,omitempty"`

	// **参数解释**: 流程的类型 - NORMAL 通用 - SURVEY 调查 - HEMOSTASIS 止血 - EASE 缓解  **约束限制**: 不涉及         **取值范围**: - NORMAL - SURVEY - HEMOSTASIS - EASE  **默认值**:  不涉及
	AopType *string `json:"aop_type,omitempty"`

	// **参数解释**: 流程的输入 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Input *string `json:"input,omitempty"`

	// **参数解释**: 流程的输出 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Output *string `json:"output,omitempty"`
}

func (o ModifyAopWorkflowVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyAopWorkflowVersionInfo struct{}"
	}

	return strings.Join([]string{"ModifyAopWorkflowVersionInfo", string(data)}, " ")
}
