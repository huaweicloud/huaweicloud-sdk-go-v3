package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClickHouseDataBaseReplicationResponse Response Object
type ListClickHouseDataBaseReplicationResponse struct {

	// 查询数据同步任务数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 数据同步任务信息。
	Replications *[]ChDatabaseReplicationInfo `json:"replications,omitempty"`

	// taurus操作表示，重启、规格变更、倒换等。
	ExtText        *string `json:"ext_text,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListClickHouseDataBaseReplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClickHouseDataBaseReplicationResponse struct{}"
	}

	return strings.Join([]string{"ListClickHouseDataBaseReplicationResponse", string(data)}, " ")
}
