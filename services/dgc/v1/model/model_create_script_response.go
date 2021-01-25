package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateScriptResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateScriptResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateScriptResponse struct{}"
	}

	return strings.Join([]string{"CreateScriptResponse", string(data)}, " ")
}
