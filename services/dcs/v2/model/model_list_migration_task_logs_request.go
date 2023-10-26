package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMigrationTaskLogsRequest Request Object
type ListMigrationTaskLogsRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 日志级别
	LogLevel *string `json:"log_level,omitempty"`
}

func (o ListMigrationTaskLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMigrationTaskLogsRequest struct{}"
	}

	return strings.Join([]string{"ListMigrationTaskLogsRequest", string(data)}, " ")
}
