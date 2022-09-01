package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMigrationTaskStatsResponse struct {

	// 全量迁移进度百分比。
	FullMigrationProgress *string `json:"full_migration_progress,omitempty" xml:"full_migration_progress"`

	// 增量迁移偏移量。
	Offset *string `json:"offset,omitempty" xml:"offset"`

	// 源redis键数量
	SourceDbsize *string `json:"source_dbsize,omitempty" xml:"source_dbsize"`

	// 目标redis键数量
	TargetDbsize *string `json:"target_dbsize,omitempty" xml:"target_dbsize"`

	// 目标redis键写入流量，单位KB/s
	TargetInputKbps *string `json:"target_input_kbps,omitempty" xml:"target_input_kbps"`

	// 目标redis每秒并发操作数
	TargetOps *string `json:"target_ops,omitempty" xml:"target_ops"`

	// 迁移任务是否在进行
	IsMigrating    *bool `json:"is_migrating,omitempty" xml:"is_migrating"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowMigrationTaskStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrationTaskStatsResponse struct{}"
	}

	return strings.Join([]string{"ShowMigrationTaskStatsResponse", string(data)}, " ")
}
