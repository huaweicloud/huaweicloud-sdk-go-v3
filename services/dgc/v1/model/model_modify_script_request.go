package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ModifyScriptRequest struct {
	ScriptName string      `json:"script_name"`
	Body       *ScriptInfo `json:"body,omitempty"`
}

func (o ModifyScriptRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ModifyScriptRequest struct{}"
	}

	return strings.Join([]string{"ModifyScriptRequest", string(data)}, " ")
}
