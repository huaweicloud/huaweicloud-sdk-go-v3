package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoUserPrivilegeV5 struct {

	// **参数解释**: 用户id。 **取值范围**: 不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**: 租户id。 **取值范围**: 不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**: 用户名。 **取值范围**: 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 仓库id。 **取值范围**: 不涉及。
	RepoId *string `json:"repo_id,omitempty"`

	// **参数解释**: 权限信息。 **取值范围**: 不涉及。
	Privilege *string `json:"privilege,omitempty"`

	// **参数解释**: 角色id。 **取值范围**: 不涉及。
	RoleId *string `json:"role_id,omitempty"`

	// **参数解释**: 角色名称。 **取值范围**: 不涉及。
	RoleName *string `json:"role_name,omitempty"`
}

func (o RepoUserPrivilegeV5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoUserPrivilegeV5 struct{}"
	}

	return strings.Join([]string{"RepoUserPrivilegeV5", string(data)}, " ")
}
