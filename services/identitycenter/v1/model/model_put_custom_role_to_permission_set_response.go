package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutCustomRoleToPermissionSetResponse Response Object
type PutCustomRoleToPermissionSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PutCustomRoleToPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutCustomRoleToPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"PutCustomRoleToPermissionSetResponse", string(data)}, " ")
}
