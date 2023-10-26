package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMigrationTaskLogsResponse Response Object
type ListMigrationTaskLogsResponse struct {

	// 日志条数
	LogNum *int32 `json:"log_num,omitempty"`

	// 日志列表
	MigrationLogs  *[]MigrationLog `json:"migration_logs,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListMigrationTaskLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMigrationTaskLogsResponse struct{}"
	}

	return strings.Join([]string{"ListMigrationTaskLogsResponse", string(data)}, " ")
}
