package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlaybookInfo Information of playbook
type PlaybookInfo struct {

	// Id value
	Id *string `json:"id,omitempty"`

	// The name, display only
	Name *string `json:"name,omitempty"`

	// The description, display only
	Description *string `json:"description,omitempty"`

	// Create time
	CreateTime *string `json:"create_time,omitempty"`

	// Update time
	UpdateTime *string `json:"update_time,omitempty"`

	// Project id value
	ProjectId *string `json:"project_id,omitempty"`

	// version Id value
	VersionId *string `json:"version_id,omitempty"`

	// If is enabled, false for disenabled, true for enabled
	Enabled *bool `json:"enabled,omitempty"`

	// 工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// Role required for approve
	ApproveRole *string `json:"approve_role,omitempty"`

	// 用户角色
	UserRole *string `json:"user_role,omitempty"`

	// Role required for edit
	EditRole *string `json:"edit_role,omitempty"`

	// Owner id
	OwnerId *string `json:"owner_id,omitempty"`

	// version
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
