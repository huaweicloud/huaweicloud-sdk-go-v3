package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IcebergSchema 定义了Iceberg表的schema结构，包含字段定义，数据类型和schema元数据。
type IcebergSchema struct {

	// 用来记录Iceberg表模式演化历史中模式版本的id。
	SchemaId int32 `json:"schema_id"`

	// 字段identifier的列表，可以识别表中的记录，用作行级操作以及去重。
	IdentifierFieldIds []int32 `json:"identifier_field_ids"`

	// 固定为struct。
	Type string `json:"type"`

	// IcebergStructField的列表。
	Fields []IcebergStructField `json:"fields"`

	// 用来指定哪些field是主键，列表内填写field_id。
	PrimaryKeys *[]int32 `json:"primary_keys,omitempty"`

	// 为该表分配的最高列ID。
	LastColumnId *int32 `json:"last_column_id,omitempty"`
}

func (o IcebergSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IcebergSchema struct{}"
	}

	return strings.Join([]string{"IcebergSchema", string(data)}, " ")
}
