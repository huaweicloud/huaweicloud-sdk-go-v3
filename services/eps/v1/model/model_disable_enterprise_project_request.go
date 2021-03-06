package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DisableEnterpriseProjectRequest struct {
	EnterpriseProjectId string `json:"enterprise_project_id"`

	Body *DisableAction `json:"body,omitempty"`
}

func (o DisableEnterpriseProjectRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisableEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"DisableEnterpriseProjectRequest", string(data)}, " ")
}
