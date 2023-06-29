package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FineGrainedSnapshotDetail 细粒度备份信息
type FineGrainedSnapshotDetail struct {

	// 数据库。
	Database *string `json:"database,omitempty"`

	// 模式列表。
	SchemaList *[]string `json:"schema_list,omitempty"`

	// 表集合。
	TableList *[]string `json:"table_list,omitempty"`
}

func (o FineGrainedSnapshotDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FineGrainedSnapshotDetail struct{}"
	}

	return strings.Join([]string{"FineGrainedSnapshotDetail", string(data)}, " ")
}
