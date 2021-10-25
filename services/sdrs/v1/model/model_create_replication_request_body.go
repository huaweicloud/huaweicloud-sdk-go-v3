package model

import (
	"encoding/json"

	"strings"
)

// 创建复制对请求体
type CreateReplicationRequestBody struct {
	Replication *CreateReplicationRequestParams `json:"replication"`
}

func (o CreateReplicationRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateReplicationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateReplicationRequestBody", string(data)}, " ")
}
