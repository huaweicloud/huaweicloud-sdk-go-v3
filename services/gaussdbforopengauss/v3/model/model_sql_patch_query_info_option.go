package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SqlPatchQueryInfoOption struct {

	// **参数解释**: unique_sql_id。 **取值范围**: 不涉及。
	SqlId *string `json:"sql_id,omitempty"`

	// **参数解释**: SQL文本。 **取值范围**: 不涉及。
	SqlText *string `json:"sql_text,omitempty"`

	// **参数解释**: SQL Patch名称。 **取值范围**: 不涉及。
	PatchName *string `json:"patch_name,omitempty"`

	// **参数解释**: 节点ID。 **取值范围**: 不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**: Patch是否生效。 **取值范围**: - true：生效 - false：不生效
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**: Patch类型。 **取值范围**: - true：Abort Patch - false：非Abort Patch
	AbortPatch *bool `json:"abort_patch,omitempty"`
}

func (o SqlPatchQueryInfoOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlPatchQueryInfoOption struct{}"
	}

	return strings.Join([]string{"SqlPatchQueryInfoOption", string(data)}, " ")
}
