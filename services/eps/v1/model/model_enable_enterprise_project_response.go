package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type EnableEnterpriseProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableEnterpriseProjectResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EnableEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"EnableEnterpriseProjectResponse", string(data)}, " ")
}
