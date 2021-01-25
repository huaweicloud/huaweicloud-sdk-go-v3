package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowScriptRequest struct {
	ScriptName string `json:"script_name"`
}

func (o ShowScriptRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowScriptRequest struct{}"
	}

	return strings.Join([]string{"ShowScriptRequest", string(data)}, " ")
}
