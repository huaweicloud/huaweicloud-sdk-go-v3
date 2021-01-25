package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAnimatedGraphicsTaskRequest struct {
	TaskId string `json:"task_id"`
}

func (o DeleteAnimatedGraphicsTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAnimatedGraphicsTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteAnimatedGraphicsTaskRequest", string(data)}, " ")
}
