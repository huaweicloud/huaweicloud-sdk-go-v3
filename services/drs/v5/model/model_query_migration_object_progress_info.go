package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryMigrationObjectProgressInfo 迁移中对象进度信息体。
type QueryMigrationObjectProgressInfo struct {

	// 概览详情。
	MigrationObjectOverview *[]MigrationObjectOverviewInfo `json:"migration_object_overview,omitempty"`

	// 数据生成时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 全量开始时间。
	FullStartTime *string `json:"full_start_time,omitempty"`

	// 全量完成时间。
	FullCompleteTime *string `json:"full_complete_time,omitempty"`

	// 增量开始时间。
	IncrStartTime *string `json:"incr_start_time,omitempty"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o QueryMigrationObjectProgressInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryMigrationObjectProgressInfo struct{}"
	}

	return strings.Join([]string{"QueryMigrationObjectProgressInfo", string(data)}, " ")
}
