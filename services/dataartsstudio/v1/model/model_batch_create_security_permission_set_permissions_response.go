package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSecurityPermissionSetPermissionsResponse Response Object
type BatchCreateSecurityPermissionSetPermissionsResponse struct {

	// 总条数
	Total *int32 `json:"total,omitempty"`

	// 权限列表
	Permissions    *[]PermissionSetPermission `json:"permissions,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o BatchCreateSecurityPermissionSetPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSecurityPermissionSetPermissionsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateSecurityPermissionSetPermissionsResponse", string(data)}, " ")
}
