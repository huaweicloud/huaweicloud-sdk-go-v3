package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListScriptsResponse struct {
	Total *int32 `json:"total,omitempty"`

	Scripts        *[]ScriptInfo `json:"scripts,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListScriptsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListScriptsResponse struct{}"
	}

	return strings.Join([]string{"ListScriptsResponse", string(data)}, " ")
}
