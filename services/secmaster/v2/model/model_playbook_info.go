package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlaybookInfo 剧本详情信息
type PlaybookInfo struct {

	// 剧本ID
	Id *string `json:"id,omitempty"`

	// 剧本名称
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 剧本创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 剧本更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 剧本版本ID
	VersionId *string `json:"version_id,omitempty"`

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 审核用户角色
	ApproveRole *string `json:"approve_role,omitempty"`

	// 用户角色
	UserRole *string `json:"user_role,omitempty"`

	// 编辑用户角色
	EditRole *string `json:"edit_role,omitempty"`

	// 所有者ID
	OwnerId *string `json:"owner_id,omitempty"`

	// 版本号
	Version *string `json:"version,omitempty"`

	// 数据类名称
	DataclassName *string `json:"dataclass_name,omitempty"`

	// 数据类ID
	DataclassId *string `json:"dataclass_id,omitempty"`

	// 待审核剧本版本ID
	UnauditedVersionId *string `json:"unaudited_version_id,omitempty"`

	// 已驳回剧本版本ID
	RejectVersionId *string `json:"reject_version_id,omitempty"`
}

func (o PlaybookInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlaybookInfo struct{}"
	}

	return strings.Join([]string{"PlaybookInfo", string(data)}, " ")
}
