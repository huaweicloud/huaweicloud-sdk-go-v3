package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GrantAuthority **参数解释**： 权限详情。 **取值范围**： 不涉及。
type GrantAuthority struct {

	// **参数解释**： 权限类型。 **取值范围**： - DATABASE：数据库。 - SCHEMA：模式。 - TABLE：表。 - VIEW：视图。 - COLUMN：列。 - FUNCTION：函数。 - SEQUENCE：序列。 - NODEGROUP：节点组。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 数据库名称。 **取值范围**： 不涉及。
	Database *string `json:"database,omitempty"`

	// **参数解释**： 模式名称。 **取值范围**： 不涉及。
	Schema *string `json:"schema,omitempty"`

	// **参数解释**： 对象名称。 **取值范围**： 不涉及。
	ObjName *string `json:"obj_name,omitempty"`

	// **参数解释**： 是否所有对象生效。 **取值范围**： 不涉及。
	AllObject *bool `json:"all_object,omitempty"`

	// **参数解释**： 是否对未来对象生效。 **取值范围**： 不涉及。
	Future *bool `json:"future,omitempty"`

	// **参数解释**： 未来对象-所属用户。 **取值范围**： 不涉及。
	FutureObjectOwners *string `json:"future_object_owners,omitempty"`

	// **参数解释**： 列名。 **取值范围**： 不涉及。
	ColumnName *[]string `json:"column_name,omitempty"`

	// **参数解释**： 权限。 **取值范围**： 不涉及。
	Privileges *[]Grant `json:"privileges,omitempty"`
}

func (o GrantAuthority) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GrantAuthority struct{}"
	}

	return strings.Join([]string{"GrantAuthority", string(data)}, " ")
}
