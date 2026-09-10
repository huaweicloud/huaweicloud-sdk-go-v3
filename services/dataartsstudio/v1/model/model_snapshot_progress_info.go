package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SnapshotProgressInfo 全量同步进度信息。
type SnapshotProgressInfo struct {

	// 全量同步整体进度。
	Progress *string `json:"progress,omitempty"`

	// 全量同步的数据库总数，如果是分库分表，则是逻辑库的数量。
	DatabaseTotal *int32 `json:"database_total,omitempty"`

	// 全量同步的数据库已读取数量，如果是分库分表，则是已读取的逻辑库数量。
	DatabaseProcessed *int32 `json:"database_processed,omitempty"`

	// 全量同步的schema总数，如果是分库分表，则是schema的数量。
	SchemaTotal *int32 `json:"schema_total,omitempty"`

	// 全量同步的schema已读取数量，如果是分库分表，则是已读取的schema的数量。
	SchemaProcessed *int32 `json:"schema_processed,omitempty"`

	// 全量同步的表总数，如果是分库分表，则是逻辑表的数量。
	TableTotal *int32 `json:"table_total,omitempty"`

	// 全量同步的表已读取数量，如果是分库分表，则是已读取的逻辑表的数量。
	TableProcessed *int32 `json:"table_processed,omitempty"`

	// 表级别全量同步进度，如果是分库分表作业，则第一层为分库分表进度。
	TableProgress *[]SnapshotTableProgressInfo `json:"table_progress,omitempty"`
}

func (o SnapshotProgressInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotProgressInfo struct{}"
	}

	return strings.Join([]string{"SnapshotProgressInfo", string(data)}, " ")
}
