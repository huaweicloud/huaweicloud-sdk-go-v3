package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListManagedRolesInPermissionSetResponse Response Object
type ListManagedRolesInPermissionSetResponse struct {

	// IAM系统策略列表
	AttachedManagedRoles *[]ResourceAttachedManagedPolicyDto `json:"attached_managed_roles,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListManagedRolesInPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManagedRolesInPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"ListManagedRolesInPermissionSetResponse", string(data)}, " ")
}
