package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateEnterpriseProjectResponse struct {
	EnterpriseProject *EpDetail `json:"enterprise_project,omitempty"`
	HttpStatusCode    int       `json:"-"`
}

func (o UpdateEnterpriseProjectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"UpdateEnterpriseProjectResponse", string(data)}, " ")
}
