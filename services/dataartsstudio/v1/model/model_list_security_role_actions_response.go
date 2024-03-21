package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityRoleActionsResponse Response Object
type ListSecurityRoleActionsResponse struct {

	// 权限操作列表
	PermissionActions *[]PermissionActions `json:"permission_actions,omitempty"`
	HttpStatusCode    int                  `json:"-"`
}

func (o ListSecurityRoleActionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityRoleActionsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityRoleActionsResponse", string(data)}, " ")
}
