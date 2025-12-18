package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseTable **参数解释：** 备份里的库表信息。 **约束限制：** 该参数仅针对GeminiDB Cassandra。 **取值范围：** - 字段为空，表示创建实例级备份。 - 字段非空，表示创建库表级备份。 **默认取值：** 不涉及。
type DatabaseTable struct {

	// **参数解释：** 数据库名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	DatabaseName string `json:"database_name"`

	// **参数解释：** 数据表名称列表。 **约束限制：** 不涉及。 **取值范围：** - table_names为空的时候，表示库级别备份。 - table_names非空的时候，表示表级别备份。  **默认取值：** 不涉及。
	TableNames *[]string `json:"table_names,omitempty"`
}

func (o DatabaseTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseTable struct{}"
	}

	return strings.Join([]string{"DatabaseTable", string(data)}, " ")
}
