package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableConfig 导入血缘请求体。
type TableConfig struct {

	// 数据库名称。
	Database *string `json:"database,omitempty"`

	// schema名称。
	Schema *string `json:"schema,omitempty"`

	// 表名称。
	Table *string `json:"table,omitempty"`

	// 是否存在。
	IsExist *bool `json:"is_exist,omitempty"`
}

func (o TableConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableConfig struct{}"
	}

	return strings.Join([]string{"TableConfig", string(data)}, " ")
}
