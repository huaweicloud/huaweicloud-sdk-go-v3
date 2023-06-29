package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePlaybookInfo Information of playbook
type CreatePlaybookInfo struct {

	// The name, display only
	Name *string `json:"name,omitempty"`

	// The description, display only
	Description *string `json:"description,omitempty"`

	// 工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// Role required for approve
	ApproveRole *string `json:"approve_role,omitempty"`

	// Role required for use
	UserRole *string `json:"user_role,omitempty"`

	// Role required for edit
	EditRole *string `json:"edit_role,omitempty"`

	// Owner id
	OwnerId *string `json:"owner_id,omitempty"`

	// If is enabled, false for disenabled, true for enabled
	Enabled *bool `json:"enabled,omitempty"`
}

func (o CreatePlaybookInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlaybookInfo struct{}"
	}

	return strings.Join([]string{"CreatePlaybookInfo", string(data)}, " ")
}
