package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CancelScriptResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelScriptResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CancelScriptResponse struct{}"
	}

	return strings.Join([]string{"CancelScriptResponse", string(data)}, " ")
}
