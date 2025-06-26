package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserAuthorityReq **参数解释**： 创建用户/角色请求体。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type UserAuthorityReq struct {

	// **参数解释**： 用户名/角色名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 类型。 **约束限制**： 不涉及。 **取值范围**： user：数据库用户。 role：数据库角色。 **默认取值**： 不涉及。
	Type string `json:"type"`

	// **参数解释**： 类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： false
	Login *bool `json:"login,omitempty"`

	// **参数解释**： 密码。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Password *string `json:"password,omitempty"`

	// **参数解释**： 是否是系统管理员。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SystemAdmin *bool `json:"system_admin,omitempty"`

	// **参数解释**： 关联的逻辑集群名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LogicalCluster *string `json:"logical_cluster,omitempty"`

	// **参数解释**： 是否允许密码登录。 **约束限制**： 不涉及。 **取值范围**： - true：允许密码登录 - false：不允许密码登录  **默认取值**： false
	PasswordDisable *bool `json:"password_disable,omitempty"`

	// **参数解释**： 是否允许创建角色。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： false
	CreateRole *bool `json:"create_role,omitempty"`

	// **参数解释**： 是否允许创建数据库。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： false
	CreateDb *bool `json:"create_db,omitempty"`

	// **参数解释**： 是否允许继承权限。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： false
	Inherit *bool `json:"inherit,omitempty"`

	// **参数解释**： 连接数限制。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ConnLimit *int32 `json:"conn_limit,omitempty"`

	// **参数解释**： 授权对象信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GrantMembers *[]string `json:"grant_members,omitempty"`

	// **参数解释**： 授权列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GrantList *[]GrantAuthority `json:"grant_list,omitempty"`

	// **参数解释**： 描述信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Desc *string `json:"desc,omitempty"`
}

func (o UserAuthorityReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserAuthorityReq struct{}"
	}

	return strings.Join([]string{"UserAuthorityReq", string(data)}, " ")
}
