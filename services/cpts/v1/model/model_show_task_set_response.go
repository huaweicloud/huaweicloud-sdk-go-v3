package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowTaskSetResponse struct {
	// code

	Code *string `json:"code,omitempty"`
	// extend

	Extend *[]string `json:"extend,omitempty"`
	// message

	Message *string `json:"message,omitempty"`
	// tasks

	Tasks          *[]ShowTaskSetResqTasks `json:"tasks,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowTaskSetResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTaskSetResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskSetResponse", string(data)}, " ")
}
