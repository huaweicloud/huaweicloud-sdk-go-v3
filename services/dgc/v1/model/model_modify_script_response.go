package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ModifyScriptResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyScriptResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ModifyScriptResponse struct{}"
	}

	return strings.Join([]string{"ModifyScriptResponse", string(data)}, " ")
}
