package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetCustomRoleForPermissionSetResponse Response Object
type GetCustomRoleForPermissionSetResponse struct {

	// 附加到权限集的自定义策略
	CustomRole     *string `json:"custom_role,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetCustomRoleForPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetCustomRoleForPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"GetCustomRoleForPermissionSetResponse", string(data)}, " ")
}
