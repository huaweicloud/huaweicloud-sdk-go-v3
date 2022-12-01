package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 迁移详情概览。
type MigrationObjectOverviewInfo struct {

	// 类型。
	Type *string `json:"type,omitempty"`

	// 待迁移数量。
	SrcCount *string `json:"src_count,omitempty"`

	// 已迁移数量。
	DstCount *string `json:"dst_count,omitempty"`

	// 状态.
	Status *string `json:"status,omitempty"`

	// 开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o MigrationObjectOverviewInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationObjectOverviewInfo struct{}"
	}

	return strings.Join([]string{"MigrationObjectOverviewInfo", string(data)}, " ")
}
