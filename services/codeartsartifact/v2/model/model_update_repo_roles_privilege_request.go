package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRepoRolesPrivilegeRequest Request Object
type UpdateRepoRolesPrivilegeRequest struct {

	// **参数解释**: 角色id。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	RoleId string `json:"role_id"`

	Body *RolePrivilegeV5 `json:"body,omitempty"`
}

func (o UpdateRepoRolesPrivilegeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepoRolesPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"UpdateRepoRolesPrivilegeRequest", string(data)}, " ")
}
