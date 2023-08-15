package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRoleResponse Response Object
type ShowRoleResponse struct {
	Role *Role `json:"role,omitempty"`

	// 角色包含的用户或者用户组信息
	UserRoles      *[]UserRole `json:"user_roles,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRoleResponse struct{}"
	}

	return strings.Join([]string{"ShowRoleResponse", string(data)}, " ")
}
