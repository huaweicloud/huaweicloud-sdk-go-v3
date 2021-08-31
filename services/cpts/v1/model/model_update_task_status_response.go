package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateTaskStatusResponse struct {
	// code

	Code *string `json:"code,omitempty"`
	// message

	Message *string `json:"message,omitempty"`
	// extend

	Extend *string `json:"extend,omitempty"`

	Result         *UpdateTaskStatusResqResult `json:"result,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o UpdateTaskStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTaskStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaskStatusResponse", string(data)}, " ")
}
