package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportTableSpaceInfo struct {

	// 数据库名。
	Database string `json:"database"`

	// 表名。
	Table string `json:"table"`

	// 数据库引擎。
	DbEngine string `json:"db_engine"`

	// 表大小。
	TableSize int64 `json:"table_size"`

	// 数据大小。
	DataSize int64 `json:"data_size"`

	// 索引大小。
	IndexSize int64 `json:"index_size"`

	// 行数量。
	Rows int64 `json:"rows"`
}

func (o HealthReportTableSpaceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportTableSpaceInfo struct{}"
	}

	return strings.Join([]string{"HealthReportTableSpaceInfo", string(data)}, " ")
}
