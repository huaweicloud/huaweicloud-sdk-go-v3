package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Column 列信息
type Column struct {

	// 是否自增
	AutoIncrement *bool `json:"auto_increment,omitempty"`

	// 数据类型
	DataType *string `json:"data_type,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 标签
	Label *string `json:"label,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 实际名称
	RealName *string `json:"real_name,omitempty"`

	// 尺寸大小
	Size *int32 `json:"size,omitempty"`

	// 表名
	TableName *string `json:"table_name,omitempty"`

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// 是否只读
	Readonly *bool `json:"readonly,omitempty"`

	// 是否二进制
	IsBinary *bool `json:"is_binary,omitempty"`

	// 数据类型
	IntDataType *int32 `json:"int_data_type,omitempty"`
}

func (o Column) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Column struct{}"
	}

	return strings.Join([]string{"Column", string(data)}, " ")
}
