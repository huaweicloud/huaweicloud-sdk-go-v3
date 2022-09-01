package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateTableResponse struct {

	// 表ID。
	TableId *string `json:"table_id,omitempty" xml:"table_id"`

	// 表名称。
	TableName *string `json:"table_name,omitempty" xml:"table_name"`

	// 表别名。
	TableAlias *string `json:"table_alias,omitempty" xml:"table_alias"`

	// 表创建时间。
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	// 表更新时间。
	ModifiedTime *string `json:"modified_time,omitempty" xml:"modified_time"`

	// 数据存储位置，分为IoTA和VIEW
	DataLocation *string `json:"data_location,omitempty" xml:"data_location"`

	// 数据类型，包括“csv”，“parquet”。
	DataType *string `json:"data_type,omitempty" xml:"data_type"`

	// 数据来源。来源类型有：pipeline, default. 默认为default.
	DataSource *string `json:"data_source,omitempty" xml:"data_source"`

	// 表类型:IoTA表为MANAGED, View为VIEW
	TableType *string `json:"table_type,omitempty" xml:"table_type"`

	// 表的描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 分区字段。只有OBS分区表有该参数，其他表没有该参数。
	PartitionColumns *[]string `json:"partition_columns,omitempty" xml:"partition_columns"`

	// 仅当数据来源为pipeline时返回。Data Store ID.
	DataStoreId *string `json:"data_store_id,omitempty" xml:"data_store_id"`

	// 标签。
	Tags           *string `json:"tags,omitempty" xml:"tags"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableResponse struct{}"
	}

	return strings.Join([]string{"CreateTableResponse", string(data)}, " ")
}
