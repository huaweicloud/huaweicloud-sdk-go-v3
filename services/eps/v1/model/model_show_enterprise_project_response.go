package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowEnterpriseProjectResponse struct {
	EnterpriseProject *EpDetail `json:"enterprise_project,omitempty"`
	HttpStatusCode    int       `json:"-"`
}

func (o ShowEnterpriseProjectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"ShowEnterpriseProjectResponse", string(data)}, " ")
}
