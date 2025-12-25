package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AopWorkflowInfo 流程详情
type AopWorkflowInfo struct {

	// **参数解释**: 流程ID **取值范围**: 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**: 流程名称 **取值范围**: 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**: 流程的描述 **取值范围**: 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**: 租户ID **取值范围**: 不涉及
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**: 所有者ID **取值范围**: 不涉及
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

	// **参数解释**: 编辑角色 **取值范围**: 不涉及
	EditRole *string `json:"edit_role,omitempty"`

	// **参数解释**: 使用角色 **取值范围**: 不涉及
	UseRole *string `json:"use_role,omitempty"`

	// **参数解释**: 审核角色 **取值范围**: 不涉及
	ApproveRole *string `json:"approve_role,omitempty"`

	// **参数解释**: 是否已启用 **取值范围**: - true  已启用 - false 未启用
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**: 流程激活版本ID **取值范围**: 不涉及
	VersionId *string `json:"version_id,omitempty"`

	// **参数解释**: 当前待审核版本ID **取值范围**: 不涉及
	CurrentApprovalVersionId *string `json:"current_approval_version_id,omitempty"`

	// **参数解释**: 当前拒绝的版本ID **取值范围**: 不涉及
	CurrentRejectedVersionId *string `json:"current_rejected_version_id,omitempty"`

	// **参数解释**: 流程的类型 - NORMAL 通用 - SURVEY 调查 - HEMOSTASIS 止血 - EASE 缓解  **取值范围**: - NORMAL - SURVEY - HEMOSTASIS - EASE
	AopType *string `json:"aop_type,omitempty"`

	// **参数解释**: 引擎的类型 - public_engine 共享版  **取值范围**: - public_engine
	EngineType *string `json:"engine_type,omitempty"`

	// **参数解释**: 数据类的ID **取值范围**: 不涉及
	DataclassId *string `json:"dataclass_id,omitempty"`

	// **参数解释**: 数据类的名称 **取值范围**: 不涉及
	DataclassName *string `json:"dataclass_name,omitempty"`

	// **参数解释**: 流程实体类型 - IP IP - ACCOUNT 账号 - DOMAIN 域名  **取值范围**: - IP - ACCOUNT - DOMAIN
	Labels *string `json:"labels,omitempty"`

	// **参数解释**: 当前激活的流程的版本 **取值范围**: 不涉及
	Version *string `json:"version,omitempty"`
}

func (o AopWorkflowInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AopWorkflowInfo struct{}"
	}

	return strings.Join([]string{"AopWorkflowInfo", string(data)}, " ")
}
