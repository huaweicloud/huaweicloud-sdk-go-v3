package model

import (
	"encoding/json"

	"strings"
)

// 复制对扩容请求体
type ExtendReplicationRequestBody struct {
	ExtendReplication *ExtendReplicationRequestParams `json:"extend-replication"`
}

func (o ExtendReplicationRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExtendReplicationRequestBody struct{}"
	}

	return strings.Join([]string{"ExtendReplicationRequestBody", string(data)}, " ")
}
