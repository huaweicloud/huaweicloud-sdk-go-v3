package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachManagedRoleFromPermissionSetResponse Response Object
type DetachManagedRoleFromPermissionSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DetachManagedRoleFromPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachManagedRoleFromPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"DetachManagedRoleFromPermissionSetResponse", string(data)}, " ")
}
