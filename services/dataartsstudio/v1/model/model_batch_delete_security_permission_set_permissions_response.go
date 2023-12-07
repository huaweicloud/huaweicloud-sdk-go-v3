package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecurityPermissionSetPermissionsResponse Response Object
type BatchDeleteSecurityPermissionSetPermissionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSecurityPermissionSetPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecurityPermissionSetPermissionsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecurityPermissionSetPermissionsResponse", string(data)}, " ")
}
