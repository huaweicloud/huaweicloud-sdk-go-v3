package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AllowDbPrivilegeRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *PostgresqlGrantRequest `json:"body,omitempty"`
}

func (o AllowDbPrivilegeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AllowDbPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"AllowDbPrivilegeRequest", string(data)}, " ")
}
