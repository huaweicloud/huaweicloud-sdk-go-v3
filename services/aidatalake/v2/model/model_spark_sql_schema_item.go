package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SparkSqlSchemaItem **参数解释**：结果数据的列结构信息，包含列名和列类型。
type SparkSqlSchemaItem struct {

	// **参数解释**：列名，用于标识结果数据中的列名称。 **取值范围**：长度为1~128个字符，例如：col1。
	ColumnName *string `json:"column_name,omitempty"`

	// **参数解释**：列类型，用于标识结果数据中的列数据类型。 **取值范围**：支持的数据类型，例如：STRING、INT等。
	ColumnType *string `json:"column_type,omitempty"`
}

func (o SparkSqlSchemaItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkSqlSchemaItem struct{}"
	}

	return strings.Join([]string{"SparkSqlSchemaItem", string(data)}, " ")
}
