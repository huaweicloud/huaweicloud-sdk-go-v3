package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ExecuteScriptResponse struct {
	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteScriptResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExecuteScriptResponse struct{}"
	}

	return strings.Join([]string{"ExecuteScriptResponse", string(data)}, " ")
}
