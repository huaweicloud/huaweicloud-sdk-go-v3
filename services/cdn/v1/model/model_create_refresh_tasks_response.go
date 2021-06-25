package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateRefreshTasksResponse struct {
	// 任务ID

	RefreshTask    *string `json:"refresh_task,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRefreshTasksResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRefreshTasksResponse struct{}"
	}

	return strings.Join([]string{"CreateRefreshTasksResponse", string(data)}, " ")
}
