package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportTableVolumeRequestBody struct {

	// **参数解释**: 数据库名称。 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及。
	DatabaseName string `json:"database_name"`

	// **参数解释**: schema名称。 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及。
	SchemaNames []string `json:"schema_names"`

	// **参数解释**:   表名称。 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及。
	TableName *string `json:"table_name,omitempty"`

	// **参数解释**: 表所属用户名称。 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 排序字段。 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及。
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释**: 排序方法。 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及。
	SortOrder *string `json:"sort_order,omitempty"`
}

func (o ExportTableVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTableVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"ExportTableVolumeRequestBody", string(data)}, " ")
}
