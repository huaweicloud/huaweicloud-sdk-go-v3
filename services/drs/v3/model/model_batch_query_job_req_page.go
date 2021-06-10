package model

import (
	"encoding/json"

	"strings"
)

// 带分页的批量查询任务详情请求体
type BatchQueryJobReqPage struct {
	// 批量查询任务详情

	Jobs []string `json:"jobs"`

	PageReq *PageReq `json:"page_req,omitempty"`
}

func (o BatchQueryJobReqPage) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchQueryJobReqPage struct{}"
	}

	return strings.Join([]string{"BatchQueryJobReqPage", string(data)}, " ")
}
