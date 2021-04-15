package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AllowSqlserverDbUserPrivilegeRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *SqlserverGrantRequest `json:"body,omitempty"`
}

func (o AllowSqlserverDbUserPrivilegeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AllowSqlserverDbUserPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"AllowSqlserverDbUserPrivilegeRequest", string(data)}, " ")
}
