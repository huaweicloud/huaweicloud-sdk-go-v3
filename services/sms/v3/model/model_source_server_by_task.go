package model

import (
	"encoding/json"

	"strings"
)

// 源端服务器信息
type SourceServerByTask struct {
	// 源端服务器id

	Id string `json:"id"`
}

func (o SourceServerByTask) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SourceServerByTask struct{}"
	}

	return strings.Join([]string{"SourceServerByTask", string(data)}, " ")
}
