package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateDatabaseRequest struct {
	// DDM实例ID

	InstanceId string `json:"instance_id"`

	Body *CreateDatabaseReq `json:"body,omitempty"`
}

func (o CreateDatabaseRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseRequest", string(data)}, " ")
}
