package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAopWorkflowVersionInfo 创建流程版本的请求信息
type CreateAopWorkflowVersionInfo struct {

	// **参数解释**: 流程名称 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Name string `json:"name"`

	// **参数解释**: 流程的描述 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Description string `json:"description"`

	// **参数解释**: 流程拓扑图的参数配置 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Taskconfig string `json:"taskconfig"`

	// **参数解释**: 流程的输入 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Input *string `json:"input,omitempty"`

	// **参数解释**: 流程的输出 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Output *string `json:"output,omitempty"`

	// **参数解释**: 流程的拓扑图的BASE64编码 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Taskflow string `json:"taskflow"`

	// **参数解释**: 流程的类型 **约束限制**: 不涉及         **取值范围**: - JSON  **默认值**:  不涉及
	TaskflowType string `json:"taskflow_type"`

	// **参数解释**: 流程的类型 - NORMAL 通用 - SURVEY 调查 - HEMOSTASIS 止血 - EASE 缓解  **约束限制**: 不涉及         **取值范围**: - NORMAL - SURVEY - HEMOSTASIS - EASE  **默认值**:  不涉及
	AopType string `json:"aop_type"`
}

func (o CreateAopWorkflowVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAopWorkflowVersionInfo struct{}"
	}

	return strings.Join([]string{"CreateAopWorkflowVersionInfo", string(data)}, " ")
}
