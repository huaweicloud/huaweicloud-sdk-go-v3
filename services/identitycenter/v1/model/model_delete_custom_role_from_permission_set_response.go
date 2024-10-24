package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomRoleFromPermissionSetResponse Response Object
type DeleteCustomRoleFromPermissionSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCustomRoleFromPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomRoleFromPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"DeleteCustomRoleFromPermissionSetResponse", string(data)}, " ")
}
