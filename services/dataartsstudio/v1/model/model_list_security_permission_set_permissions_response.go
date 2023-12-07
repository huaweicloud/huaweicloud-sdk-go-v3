package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityPermissionSetPermissionsResponse Response Object
type ListSecurityPermissionSetPermissionsResponse struct {

	// 总条数
	Total *int32 `json:"total,omitempty"`

	// 权限列表
	Permissions    *[]PermissionSetPermission `json:"permissions,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListSecurityPermissionSetPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityPermissionSetPermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityPermissionSetPermissionsResponse", string(data)}, " ")
}
