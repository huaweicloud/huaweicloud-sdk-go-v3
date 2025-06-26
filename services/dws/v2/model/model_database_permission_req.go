package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabasePermissionReq **参数解释**： 数据库权限请求。 **取值范围**： 不涉及。
type DatabasePermissionReq struct {

	// **参数解释**： 对象类型。 **取值范围**： DATABASE、SCHEMA、TABLE、VIEW、COLUMN、FUNCTION、SEQUENCE、NODEGROUP、ROLE。
	Type string `json:"type"`

	// **参数解释**： 是否授权操作。 **取值范围**： 不涉及。
	IsGrant bool `json:"is_grant"`

	// **参数解释**： 授权列表。is_grant为true时必填。 **取值范围**： 不涉及。
	GrantList *[]Grant `json:"grant_list,omitempty"`

	// **参数解释**： 撤销权限列表。is_grant为false时必填。 **取值范围**： 不涉及。
	RevokeList *[]Revoke `json:"revoke_list,omitempty"`

	// **参数解释**： 被授权角色列表。 **取值范围**： 不涉及。
	RoleList []string `json:"role_list"`

	// **参数解释**： 权限所属对象列表。 **取值范围**： 不涉及。
	ObjectList []string `json:"object_list"`

	// **参数解释**： schema下所有数据库对象权限，默认false。 **取值范围**： 不涉及。
	AllObject *bool `json:"all_object,omitempty"`

	// **参数解释**： 撤销权限是否级联撤销，默认true。 **取值范围**： 不涉及。
	Cascade *bool `json:"cascade,omitempty"`

	// **参数解释**： 数据库名称。 **取值范围**： 不涉及。
	Database string `json:"database"`

	// **参数解释**： 模式名称。 **取值范围**： 不涉及。
	Schema *string `json:"schema,omitempty"`

	// **参数解释**： 表名。 **取值范围**： 不涉及。
	Table *string `json:"table,omitempty"`
}

func (o DatabasePermissionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabasePermissionReq struct{}"
	}

	return strings.Join([]string{"DatabasePermissionReq", string(data)}, " ")
}
