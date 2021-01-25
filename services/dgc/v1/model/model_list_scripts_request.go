package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListScriptsRequest struct {
}

func (o ListScriptsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListScriptsRequest struct{}"
	}

	return strings.Join([]string{"ListScriptsRequest", string(data)}, " ")
}
