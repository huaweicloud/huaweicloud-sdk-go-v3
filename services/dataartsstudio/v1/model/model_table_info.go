package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableInfo struct {

	// 数据库名称
	Database *string `json:"database,omitempty"`

	// 逻辑库名称
	Schema *string `json:"schema,omitempty"`

	// 表名称
	Table *string `json:"table,omitempty"`
}

func (o TableInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableInfo struct{}"
	}

	return strings.Join([]string{"TableInfo", string(data)}, " ")
}
