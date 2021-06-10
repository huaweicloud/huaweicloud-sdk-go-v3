package model

import (
	"encoding/json"

	"strings"
)

// 批量修改任务请求体
type BatchModifyJobReq struct {
	// 修改任务请求体

	Jobs []ModifyJobReq `json:"jobs"`
}

func (o BatchModifyJobReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchModifyJobReq struct{}"
	}

	return strings.Join([]string{"BatchModifyJobReq", string(data)}, " ")
}
