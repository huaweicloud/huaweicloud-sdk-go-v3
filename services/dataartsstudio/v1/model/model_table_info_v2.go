package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TableInfoV2 struct {

	// 目录名称
	Catalog *string `json:"catalog,omitempty"`

	// 数据库名称
	Database *string `json:"database,omitempty"`

	// 逻辑库名称
	Schema *string `json:"schema,omitempty"`

	// 表名称
	Table *string `json:"table,omitempty"`
}

func (o TableInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableInfoV2 struct{}"
	}

	return strings.Join([]string{"TableInfoV2", string(data)}, " ")
}
