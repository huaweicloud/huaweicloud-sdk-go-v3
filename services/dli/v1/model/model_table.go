package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Table 查询所有表响应参数的表的信息。 说明： 若URI中的过滤条件“with-detail=false”，则tables相关参数中只返回“data_location”，“table_name”，“table_type”三个参数。
type Table struct {

	// 表创建时间。是单位为“毫秒”的时间戳。
	CreateTime int64 `json:"create_time"`

	// 所列OBS表数据的类型，目前支持：parquet、ORC、CSV、JSON、Carbon、Avro格式。只有OBS表有该参数，其他表没有该参数。
	DataType *string `json:"data_type,omitempty"`

	// 数据存储位置，分为MANAGED，EXTERNAL和VIEW，其中EXTERNAL包括OBS表、CLoudTable表。
	DataLocation string `json:"data_location"`

	// 最近更新时间。是单位为“毫秒”的时间戳。
	LastAccessTime int64 `json:"last_access_time"`

	// OBS表上的存储路径。
	Location *string `json:"location,omitempty"`

	// 表创建者。
	Owner string `json:"owner"`

	// 表名称。
	TableName string `json:"table_name"`

	// DLI表的大小。非DLI表该参数值为0。
	TableSize int64 `json:"table_size"`

	// 表类型： OBS表为EXTERNAL DLI表为MANAGED View为VIEW
	TableType string `json:"table_type"`

	// 分区字段。只有OBS分区表有该参数，其他表没有该参数。
	PartitionColumns *[]string `json:"partition_columns,omitempty"`

	// 分页大小，最小为1，最大为100。
	PageSize *int32 `json:"page-size,omitempty"`

	// 当前页码，最小为1。
	CurrentPage *int32 `json:"current-page,omitempty"`
}

func (o Table) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Table struct{}"
	}

	return strings.Join([]string{"Table", string(data)}, " ")
}
