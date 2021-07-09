package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RebuildConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RebuildConfigResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RebuildConfigResponse struct{}"
	}

	return strings.Join([]string{"RebuildConfigResponse", string(data)}, " ")
}
