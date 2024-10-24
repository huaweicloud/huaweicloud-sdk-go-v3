package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachManagedRoleToPermissionSetResponse Response Object
type AttachManagedRoleToPermissionSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AttachManagedRoleToPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachManagedRoleToPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"AttachManagedRoleToPermissionSetResponse", string(data)}, " ")
}
