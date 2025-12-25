package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatabaseUserResponse Response Object
type ShowDatabaseUserResponse struct {

	// **参数解释**： 用户名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 是否可以登录。 **取值范围**： 不涉及。
	Login *bool `json:"login,omitempty"`

	// **参数解释**： 创建角色权限。 **取值范围**： 不涉及。
	Createrole *bool `json:"createrole,omitempty"`

	// **参数解释**： 创建数据库权限。 **取值范围**： 不涉及。
	Createdb *bool `json:"createdb,omitempty"`

	// **参数解释**： 系统管理员。 **取值范围**： 不涉及。
	Systemadmin *bool `json:"systemadmin,omitempty"`

	// **参数解释**： 审计管理员。 **取值范围**： 不涉及。
	Auditadmin *bool `json:"auditadmin,omitempty"`

	// **参数解释**： 继承所在组权限。 **取值范围**： 不涉及。
	Inherit *bool `json:"inherit,omitempty"`

	// **参数解释**： 访问外表权限。 **取值范围**： 不涉及。
	Useft *bool `json:"useft,omitempty"`

	// **参数解释**： 连接数限制。 **取值范围**： 不涉及。
	ConnLimit *int32 `json:"conn_limit,omitempty"`

	// **参数解释**： 是否允许流复制。 **取值范围**： 不涉及。
	Replication *bool `json:"replication,omitempty"`

	// **参数解释**： 角色生效时间。 **取值范围**： 不涉及。
	ValidBegin *int64 `json:"valid_begin,omitempty"`

	// **参数解释**： 角色过期时间。 **取值范围**： 不涉及。
	ValidUntil *int64 `json:"valid_until,omitempty"`

	// **参数解释**： 是否锁定。 **取值范围**： 不涉及。
	Lock *bool `json:"lock,omitempty"`

	// **参数解释**： 描述。 **取值范围**： 不涉及。
	Desc *string `json:"desc,omitempty"`

	// **参数解释**： 用户类型。 **取值范围**： - COMMON：公共。
	UserType *string `json:"user_type,omitempty"`

	// **参数解释**： description: 所属逻辑集群。 **取值范围**： 不涉及。
	LogicalCluster *string `json:"logical_cluster,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"ShowDatabaseUserResponse", string(data)}, " ")
}
