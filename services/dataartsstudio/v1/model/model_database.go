package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Database struct {

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 数据库guid
	DatabaseGuid *string `json:"database_guid,omitempty"`

	// 数据库的唯一标识名称
	DatabaseQualifiedName *string `json:"database_qualified_name,omitempty"`

	// 数据库中表数目
	TableCount *int32 `json:"table_count,omitempty"`

	// 数据量大小
	DataSize *float64 `json:"data_size,omitempty"`
}

func (o Database) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Database struct{}"
	}

	return strings.Join([]string{"Database", string(data)}, " ")
}
