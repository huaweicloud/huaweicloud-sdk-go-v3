package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteDatabaseResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDatabaseResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDatabaseResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseResponse", string(data)}, " ")
}
