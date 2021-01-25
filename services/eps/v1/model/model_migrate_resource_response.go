package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type MigrateResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o MigrateResourceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MigrateResourceResponse struct{}"
	}

	return strings.Join([]string{"MigrateResourceResponse", string(data)}, " ")
}
