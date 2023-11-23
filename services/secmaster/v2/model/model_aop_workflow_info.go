package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AopWorkflowInfo 流程详情
type AopWorkflowInfo struct {

	// 流程ID
	Id *string `json:"id,omitempty"`

	// 流程名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 租户ID
	ProjectId *string `json:"project_id,omitempty"`

	// 所有者ID
	OwnerId *string `json:"owner_id,omitempty"`

	// 创建者ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 编辑角色
	EditRole *string `json:"edit_role,omitempty"`

	// 是用角色
	UseRole *string `json:"use_role,omitempty"`

	// 审核人
	ApproveRole *string `json:"approve_role,omitempty"`

	// 是否已启用
	Enabled *bool `json:"enabled,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 流程版本ID
	VersionId *string `json:"version_id,omitempty"`

	// 当前待审核版本号
	CurrentApprovalVersionId *string `json:"current_approval_version_id,omitempty"`

	// 当前拒绝的版本号
	CurrentRejectedVersoinId *string `json:"current_rejected_versoin_id,omitempty"`

	// aop的类型有以下的值     NORMAL, 通用     SURVEY, 调查     HEMOSTASIS,止血     EASE;缓解
	AopType *string `json:"aop_type,omitempty"`

	// 引擎的类型分为共享版和专项版
	EngineType *string `json:"engine_type,omitempty"`

	// 数据类的ID
	DataclassId *string `json:"dataclass_id,omitempty"`
}

func (o AopWorkflowInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AopWorkflowInfo struct{}"
	}

	return strings.Join([]string{"AopWorkflowInfo", string(data)}, " ")
}
