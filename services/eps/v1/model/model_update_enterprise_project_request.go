package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateEnterpriseProjectRequest struct {
	EnterpriseProjectId string             `json:"enterprise_project_id"`
	Body                *EnterpriseProject `json:"body,omitempty"`
}

func (o UpdateEnterpriseProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"UpdateEnterpriseProjectRequest", string(data)}, " ")
}
