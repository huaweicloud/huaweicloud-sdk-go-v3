package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowNewTaskStatusResponse struct {
	// 任务状态，有以下几种： - success - failed - waiting - running - preprocess - ready

	TaskStatus *string `json:"task_status,omitempty"`
	// 任务的附加信息

	TaskMsg        *string `json:"task_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowNewTaskStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowNewTaskStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowNewTaskStatusResponse", string(data)}, " ")
}
