package model

import (
	"encoding/json"

	"strings"
)

// 批量预检查请求体
type BatchPrecheckReq struct {
	// 批量预检查请求列表

	Jobs []PreCheckInfo `json:"jobs"`
}

func (o BatchPrecheckReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchPrecheckReq struct{}"
	}

	return strings.Join([]string{"BatchPrecheckReq", string(data)}, " ")
}
