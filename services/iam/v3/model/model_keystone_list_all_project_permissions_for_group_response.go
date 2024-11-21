package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListAllProjectPermissionsForGroupResponse Response Object
type KeystoneListAllProjectPermissionsForGroupResponse struct {
	Links *Links `json:"links,omitempty"`

	// 权限信息列表。
	Roles          *[]InheritedRoleResult `json:"roles,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o KeystoneListAllProjectPermissionsForGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListAllProjectPermissionsForGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListAllProjectPermissionsForGroupResponse", string(data)}, " ")
}
