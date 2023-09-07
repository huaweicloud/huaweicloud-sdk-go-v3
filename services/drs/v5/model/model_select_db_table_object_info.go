package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SelectDbTableObjectInfo 指定库表信息
type SelectDbTableObjectInfo struct {

	// 数据库名称。
	DbName string `json:"db_name"`

	// 数据库schema名称。
	SchemaName *string `json:"schema_name,omitempty"`

	// 数据库表名称。
	TableName string `json:"table_name"`
}

func (o SelectDbTableObjectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SelectDbTableObjectInfo struct{}"
	}

	return strings.Join([]string{"SelectDbTableObjectInfo", string(data)}, " ")
}
