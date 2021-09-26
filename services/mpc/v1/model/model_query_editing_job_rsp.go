package model

import (
	"encoding/json"

	"strings"
)

type QueryEditingJobRsp struct {
	// 任务总数

	Total *int32 `json:"total,omitempty"`
	// 任务列表

	Jobs *[]EditingJob `json:"jobs,omitempty"`
}

func (o QueryEditingJobRsp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryEditingJobRsp struct{}"
	}

	return strings.Join([]string{"QueryEditingJobRsp", string(data)}, " ")
}
