package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTableResponse Response Object
type ShowTableResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 表的总列数。
	ColumnCount *int32 `json:"column_count,omitempty"`

	// 表的列信息，包含列名称、类型和描述信息。
	Columns *[]Column `json:"columns,omitempty"`

	// 表类型，包括“MANAGED”，“EXTERNAL”和“VIEW”。
	TableType *string `json:"table_type,omitempty"`

	// 数据类型，包括“csv”，“parquet”，“orc”，“json”，“carbon”和“avro”。
	DataType *string `json:"data_type,omitempty"`

	// 数据存储的路径，以“s3a”开头。
	DataLocation *string `json:"data_location,omitempty"`

	// 存储属性，以“key/value”的格式出现，包含delimiter，escape，quote，header，dateformat，timestampformat参数。
	StorageProperties *[]interface{} `json:"storage_properties,omitempty"`

	// 表的注释。
	TableComment *string `json:"table_comment,omitempty"`

	// 建表的语句
	CreateTableSql *string `json:"create_table_sql,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableResponse struct{}"
	}

	return strings.Join([]string{"ShowTableResponse", string(data)}, " ")
}
