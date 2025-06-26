package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseUserInfoReq **参数解释**： 用户详细信息。 **取值范围**： 不涉及。
type DatabaseUserInfoReq struct {

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

	// **参数解释**： 角色生效时间。格式：yyyy-MM-ddTHH:mm:ssZ。 **取值范围**： 不涉及。
	ValidBegin *string `json:"valid_begin,omitempty"`

	// **参数解释**： 角色过期时间。格式：yyyy-MM-ddTHH:mm:ssZ。 **取值范围**： 不涉及。
	ValidUntil *string `json:"valid_until,omitempty"`

	// **参数解释**： 是否锁定。 **取值范围**： 不涉及。
	Lock *bool `json:"lock,omitempty"`
}

func (o DatabaseUserInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseUserInfoReq struct{}"
	}

	return strings.Join([]string{"DatabaseUserInfoReq", string(data)}, " ")
}
