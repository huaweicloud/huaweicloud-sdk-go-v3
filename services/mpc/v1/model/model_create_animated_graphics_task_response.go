package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateAnimatedGraphicsTaskResponse struct {
	// 任务ID

	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAnimatedGraphicsTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAnimatedGraphicsTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateAnimatedGraphicsTaskResponse", string(data)}, " ")
}
