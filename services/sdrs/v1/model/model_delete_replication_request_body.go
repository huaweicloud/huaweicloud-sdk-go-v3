package model

import (
	"encoding/json"

	"strings"
)

// 删除复制对请求体
type DeleteReplicationRequestBody struct {
	Replication *DeleteReplicationRequestParams `json:"replication"`
}

func (o DeleteReplicationRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteReplicationRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteReplicationRequestBody", string(data)}, " ")
}
