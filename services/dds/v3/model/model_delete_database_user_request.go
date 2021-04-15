package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteDatabaseUserRequest struct {
	InstanceId string `json:"instance_id"`

	Body *DeleteDatabaseUserRequestBody `json:"body,omitempty"`
}

func (o DeleteDatabaseUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseUserRequest", string(data)}, " ")
}
