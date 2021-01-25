package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteScriptResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScriptResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteScriptResponse struct{}"
	}

	return strings.Join([]string{"DeleteScriptResponse", string(data)}, " ")
}
