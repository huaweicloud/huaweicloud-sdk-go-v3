package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AopWorkflowVersionInfo 流程版本详情信息
type AopWorkflowVersionInfo struct {

	// **参数解释**: 流程版本ID **约束限制**: 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**: 流程版本ID **约束限制**: 不涉及
	VersionId *string `json:"version_id,omitempty"`

	// **参数解释**: 流程版本名字 **约束限制**: 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**: 流程ID **约束限制**: 不涉及
	AopworkflowId *string `json:"aopworkflow_id,omitempty"`

	// **参数解释**: 流程版本的描述 **取值范围**: 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**: 租户的项目ID **约束限制**: 不涉及
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**: 所有者ID **约束限制**: 不涉及
	OwnerId *string `json:"owner_id,omitempty"`

	// **参数解释**: 创建者ID **取值范围**: 不涉及
	CreatorId *string `json:"creator_id,omitempty"`

	// **参数解释**: 创建者的名称 **取值范围**: 不涉及
	CreatorName *string `json:"creator_name,omitempty"`

	// **参数解释**: 创建时间 **取值范围**: 不涉及
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释**: 更新者的ID **取值范围**: 不涉及
	ModifierId *string `json:"modifier_id,omitempty"`

	// **参数解释**: 更新者的名称 **取值范围**: 不涉及
	ModifierName *string `json:"modifier_name,omitempty"`

	// **参数解释**: 更新的时间 **取值范围**: 不涉及
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**: 是否已启用 **约束限制**: 不涉及         **取值范围**: - true  已启用 - false 未启用  **默认值**:  false
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**: 流程的状态 - pending_submit 草稿 - pending_approval 待审核 - publishing 发布中 - publish_failed 发布失败 - not_activated 未激活 - activated 已激活 - rejected 审核拒绝  **取值范围**: - pending_submit - pending_approval - publishing - publish_failed - not_activated - activated - rejected
	Status *string `json:"status,omitempty"`

	// **参数解释**: 当前流程的版本 **约束限制**: 不涉及         **取值范围**: 不涉及 **默认值**:  不涉及
	Version *string `json:"version,omitempty"`

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

	// **参数解释**: 数据类的ID **取值范围**: 不涉及
	DataclassId *string `json:"dataclass_id,omitempty"`

	// **参数解释**: 数据类的名称 **取值范围**: 不涉及
	DataclassName *string `json:"dataclass_name,omitempty"`

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o AopWorkflowVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AopWorkflowVersionInfo struct{}"
	}

	return strings.Join([]string{"AopWorkflowVersionInfo", string(data)}, " ")
}
