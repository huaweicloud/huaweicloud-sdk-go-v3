package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListTakeOverTaskResponse struct {
	Tasks *[]TakeOverTask `json:"tasks,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTakeOverTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTakeOverTaskResponse struct{}"
	}

	return strings.Join([]string{"ListTakeOverTaskResponse", string(data)}, " ")
}
