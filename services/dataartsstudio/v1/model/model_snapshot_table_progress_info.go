package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SnapshotTableProgressInfo 表级别全量同步进度。
type SnapshotTableProgressInfo struct {

	// 单表全量同步进度。
	Progress *string `json:"progress,omitempty"`

	// 数据库名或者逻辑库名。
	Database *string `json:"database,omitempty"`

	// 表名或者逻辑表名。
	Table *string `json:"table,omitempty"`

	// 读取的表的数据条数。
	RecordsRead *int64 `json:"records_read,omitempty"`

	// 是否是逻辑表。
	LogicTable *bool `json:"logic_table,omitempty"`

	// 表级别全量同步进度，如果是分库分表作业，则这里为各个子表的同步进度。
	SubTableProgress *[]SnapshotTableProgressInfo `json:"sub_table_progress,omitempty"`
}

func (o SnapshotTableProgressInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotTableProgressInfo struct{}"
	}

	return strings.Join([]string{"SnapshotTableProgressInfo", string(data)}, " ")
}
