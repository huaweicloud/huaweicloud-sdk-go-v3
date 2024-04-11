package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplayingSqlResp 正在回放的SQL信息
type ReplayingSqlResp struct {

	// 库名或shema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// SQL语句
	SqlStatement *string `json:"sql_statement,omitempty"`

	// 原始执行耗时
	Latency *int32 `json:"latency,omitempty"`

	// 回放执行耗时
	ExecuteLatency *int32 `json:"execute_latency,omitempty"`

	// 执行状态
	Status *string `json:"status,omitempty"`

	// 客户端IP
	Client *string `json:"client,omitempty"`

	// 连接ID
	ConnectionId *string `json:"connection_id,omitempty"`

	// 回放开始时间
	ReplayStartTime *string `json:"replay_start_time,omitempty"`
}

func (o ReplayingSqlResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplayingSqlResp struct{}"
	}

	return strings.Join([]string{"ReplayingSqlResp", string(data)}, " ")
}
