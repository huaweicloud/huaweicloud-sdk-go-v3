package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteDatabaseRoleRequest struct {
	InstanceId string `json:"instance_id"`

	Body *DeleteDatabaseRoleRequestBody `json:"body,omitempty"`
}

func (o DeleteDatabaseRoleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDatabaseRoleRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseRoleRequest", string(data)}, " ")
}
