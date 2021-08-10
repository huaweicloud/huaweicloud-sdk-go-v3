package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowJobListByProjectIdResponse struct {
	// 任务列表

	Jobs *[]Job `json:"jobs,omitempty"`
	// 任务总数

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowJobListByProjectIdResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobListByProjectIdResponse struct{}"
	}

	return strings.Join([]string{"ShowJobListByProjectIdResponse", string(data)}, " ")
}
