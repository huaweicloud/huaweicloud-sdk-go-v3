package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTableSchemaResponse struct {

	// 表ID。
	TableId *string `json:"table_id,omitempty" xml:"table_id"`

	// 表名。
	TableName *string `json:"table_name,omitempty" xml:"table_name"`

	// 表别名。
	TableAlias *string `json:"table_alias,omitempty" xml:"table_alias"`

	// 标签。
	Tags *string `json:"tags,omitempty" xml:"tags"`

	// 表的总列数。
	ColumnCount *int32 `json:"column_count,omitempty" xml:"column_count"`

	// 表的列信息，包含列名称、类型和描述信息。
	Columns *[]Column `json:"columns,omitempty" xml:"columns"`

	// 表类型，包括“MANAGED”，“EXTERNAL”和“VIEW”。
	TableType *string `json:"table_type,omitempty" xml:"table_type"`

	// 数据类型，包括“csv”，“parquet”。
	DataType *string `json:"data_type,omitempty" xml:"data_type"`

	// 数据存储的路径，以“s3a”开头。
	DataLocation *string `json:"data_location,omitempty" xml:"data_location"`

	StorageProperties *[]KeyValue `json:"storage_properties,omitempty" xml:"storage_properties"`
	HttpStatusCode    int         `json:"-"`
}

func (o ShowTableSchemaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableSchemaResponse struct{}"
	}

	return strings.Join([]string{"ShowTableSchemaResponse", string(data)}, " ")
}
