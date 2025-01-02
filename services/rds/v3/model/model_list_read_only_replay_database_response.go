package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReadOnlyReplayDatabaseResponse Response Object
type ListReadOnlyReplayDatabaseResponse struct {

	// 每次返回的库上限数量
	DatabaseLimit *int32 `json:"database_limit,omitempty"`

	// 返回的总表数量
	TotalTables *int32 `json:"total_tables,omitempty"`

	// 每次返回的表上限数量
	TableLimit *int32 `json:"table_limit,omitempty"`

	// 可恢复到主实例的数据库列表
	Databases      *[]DelayRestoreDatabase `json:"databases,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListReadOnlyReplayDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReadOnlyReplayDatabaseResponse struct{}"
	}

	return strings.Join([]string{"ListReadOnlyReplayDatabaseResponse", string(data)}, " ")
}
