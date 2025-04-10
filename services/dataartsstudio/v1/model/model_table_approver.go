package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableApprover 审批人
type TableApprover struct {

	// 审批人id
	Id *string `json:"id,omitempty"`

	// 审批人姓名
	Name *string `json:"name,omitempty"`

	// 权限集id
	PermissionSetId *string `json:"permission_set_id,omitempty"`

	// 审批人类型, 0 用户  1 用户组  2 角色（空间管理员）
	Type *int32 `json:"type,omitempty"`

	// 空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 安全管理员：SECURITY_MANAGER，空间权限集管理员WORKSPACE_PERMISSION_SET_MANAGER，权限集管理员PERMISSION_SET_MANAGER
	TableApproverType *string `json:"table_approver_type,omitempty"`
}

func (o TableApprover) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableApprover struct{}"
	}

	return strings.Join([]string{"TableApprover", string(data)}, " ")
}
