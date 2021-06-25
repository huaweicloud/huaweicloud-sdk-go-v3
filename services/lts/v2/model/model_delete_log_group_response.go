package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteLogGroupResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteLogGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteLogGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteLogGroupResponse", string(data)}, " ")
}
