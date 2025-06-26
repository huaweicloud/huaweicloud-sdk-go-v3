package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTablesStatisticDto **参数解释**： 表倾斜率或脏页率列表。 **取值范围**： 不涉及。
type ListTablesStatisticDto struct {

	// **参数解释**： 数据库名称。 **取值范围**： 不涉及。
	DbName *string `json:"db_name,omitempty"`

	// **参数解释**： 模式名称。 **取值范围**： 不涉及。
	SchemaName *string `json:"schema_name,omitempty"`

	// **参数解释**： 表名。 **取值范围**： 不涉及。
	TableName *string `json:"table_name,omitempty"`

	// **参数解释**： 所属用户。 **取值范围**： 不涉及。
	TableOwner *string `json:"table_owner,omitempty"`

	// **参数解释**： 表大小。 **取值范围**： 不涉及。
	TableSize *string `json:"table_size,omitempty"`

	// **参数解释**： 表倾斜率。 **取值范围**： 不涉及。
	SkewRate *float64 `json:"skew_rate,omitempty"`

	// **参数解释**： 脏页率。 **取值范围**： 不涉及。
	DirtyPageRate *float64 `json:"dirty_page_rate,omitempty"`
}

func (o ListTablesStatisticDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablesStatisticDto struct{}"
	}

	return strings.Join([]string{"ListTablesStatisticDto", string(data)}, " ")
}
