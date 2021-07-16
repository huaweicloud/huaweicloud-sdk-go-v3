package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateRedislogRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 日期偏移量，表示从过去的n天开始查询，例如：传入0则表示查询今天的日志，传入7则表示查询过去7天的日志。最大支持0-7。

	QueryTime *int32 `json:"query_time,omitempty"`
	// 返回日志的类型，当前仅支持Redis运行日志，类型为run

	LogType string `json:"log_type"`
	// 副本ID，可以从分片与副本中查询对应节点的副本ID

	ReplicationId *string `json:"replication_id,omitempty"`
}

func (o CreateRedislogRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRedislogRequest struct{}"
	}

	return strings.Join([]string{"CreateRedislogRequest", string(data)}, " ")
}
