package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreatePreheatingTasksResponse struct {
	// 任务ID

	PreheatingTask *string `json:"preheating_task,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePreheatingTasksResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePreheatingTasksResponse struct{}"
	}

	return strings.Join([]string{"CreatePreheatingTasksResponse", string(data)}, " ")
}
