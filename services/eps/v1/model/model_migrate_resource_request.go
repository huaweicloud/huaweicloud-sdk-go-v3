package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type MigrateResourceRequest struct {
	EnterpriseProjectId string `json:"enterprise_project_id"`

	Body *MigrateResource `json:"body,omitempty"`
}

func (o MigrateResourceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MigrateResourceRequest struct{}"
	}

	return strings.Join([]string{"MigrateResourceRequest", string(data)}, " ")
}
