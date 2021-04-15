package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteDatabaseRoleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDatabaseRoleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDatabaseRoleResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseRoleResponse", string(data)}, " ")
}
