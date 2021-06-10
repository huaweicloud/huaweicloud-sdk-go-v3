package model

import (
	"encoding/json"

	"strings"
)

// 批量结束与删除在线迁移任务请求体
type BatchDeleteJobReq struct {
	// 批量结束与删除任务请求列表

	Jobs []DeleteJobReq `json:"jobs"`
}

func (o BatchDeleteJobReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteJobReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobReq", string(data)}, " ")
}
