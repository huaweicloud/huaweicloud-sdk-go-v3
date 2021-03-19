package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListBatchTasksResponse struct {
	// 批量任务列表。

	Batchtasks *[]Task `json:"batchtasks,omitempty"`

	Page           *Page `json:"page,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListBatchTasksResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBatchTasksResponse struct{}"
	}

	return strings.Join([]string{"ListBatchTasksResponse", string(data)}, " ")
}
