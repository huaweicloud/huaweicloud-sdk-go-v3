package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteMigrationTaskRequest struct {
	Body *DeleteMigrateTaskRequest `json:"body,omitempty"`
}

func (o DeleteMigrationTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteMigrationTaskRequest", string(data)}, " ")
}
