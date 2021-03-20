package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListScriptResultsRequest struct {
	ScriptName string `json:"script_name"`

	InstanceId string `json:"instance_id"`
}

func (o ListScriptResultsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListScriptResultsRequest struct{}"
	}

	return strings.Join([]string{"ListScriptResultsRequest", string(data)}, " ")
}
