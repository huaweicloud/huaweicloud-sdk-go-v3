package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ColumnConfig 导入血缘请求体。
type ColumnConfig struct {

	// 数据库名称。
	Database string `json:"database"`

	// schema名称。
	Schema *string `json:"schema,omitempty"`

	// 表名称。
	Table string `json:"table"`

	// 字段名称。
	Column string `json:"column"`

	// 是否存在。
	IsExist *bool `json:"is_exist,omitempty"`
}

func (o ColumnConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ColumnConfig struct{}"
	}

	return strings.Join([]string{"ColumnConfig", string(data)}, " ")
}
