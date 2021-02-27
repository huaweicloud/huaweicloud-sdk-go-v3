package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateScriptResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateScriptResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateScriptResponse struct{}"
	}

	return strings.Join([]string{"UpdateScriptResponse", string(data)}, " ")
}
