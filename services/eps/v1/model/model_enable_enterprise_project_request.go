package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type EnableEnterpriseProjectRequest struct {
	EnterpriseProjectId string `json:"enterprise_project_id"`

	Body *EnableAction `json:"body,omitempty"`
}

func (o EnableEnterpriseProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EnableEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"EnableEnterpriseProjectRequest", string(data)}, " ")
}
