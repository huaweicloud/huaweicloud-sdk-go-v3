package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowTaskResponse struct {
	// code

	Code *string `json:"code,omitempty"`
	// message

	Message *string `json:"message,omitempty"`

	Taskinfo       *ShowTaskResqTaskinfo `json:"taskinfo,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskResponse", string(data)}, " ")
}
