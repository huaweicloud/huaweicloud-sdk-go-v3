package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityUnreasonablePermissionsResponse Response Object
type ListSecurityUnreasonablePermissionsResponse struct {

	// 不合理权限配置总条数。
	Total *int32 `json:"total,omitempty"`

	// 不合理权限配置列表。
	UnreasonablePermissions *[]DiagnosePermissionDetail `json:"unreasonable_permissions,omitempty"`
	HttpStatusCode          int                         `json:"-"`
}

func (o ListSecurityUnreasonablePermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityUnreasonablePermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityUnreasonablePermissionsResponse", string(data)}, " ")
}
