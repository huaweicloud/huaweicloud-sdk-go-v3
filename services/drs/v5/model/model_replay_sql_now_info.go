package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplaySqlNowInfo 正在回放的sql信息
type ReplaySqlNowInfo struct {

	// 会话id
	ThreadId int64 `json:"thread_id"`

	// 创建时间
	CreatedAt int64 `json:"created_at"`

	// 修改时间
	ModifiedAt int64 `json:"modified_at"`

	// 分片id
	ShardId int64 `json:"shard_id"`

	// schema名称
	SchemaName string `json:"schema_name"`

	// sql语句
	SqlStatement string `json:"sql_statement"`

	// 原始耗时
	Latency int64 `json:"latency"`

	// 执行耗时
	ExecuteLatency int64 `json:"execute_latency"`

	// 目标库类型
	TargetType string `json:"target_type"`

	// 目标库标识
	TargetName string `json:"target_name"`

	// 回放状态。取值： - running：执行中。 - finish：已完成。
	Status string `json:"status"`
}

func (o ReplaySqlNowInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplaySqlNowInfo struct{}"
	}

	return strings.Join([]string{"ReplaySqlNowInfo", string(data)}, " ")
}
