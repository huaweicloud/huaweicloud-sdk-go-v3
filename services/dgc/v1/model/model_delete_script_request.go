package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteScriptRequest struct {
	ScriptName string `json:"script_name"`
}

func (o DeleteScriptRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteScriptRequest struct{}"
	}

	return strings.Join([]string{"DeleteScriptRequest", string(data)}, " ")
}
