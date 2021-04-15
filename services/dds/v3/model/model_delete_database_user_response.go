package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteDatabaseUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDatabaseUserResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseUserResponse", string(data)}, " ")
}
