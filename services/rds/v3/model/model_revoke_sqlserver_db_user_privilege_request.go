package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RevokeSqlserverDbUserPrivilegeRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *SqlserverRevokeRequest `json:"body,omitempty"`
}

func (o RevokeSqlserverDbUserPrivilegeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RevokeSqlserverDbUserPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"RevokeSqlserverDbUserPrivilegeRequest", string(data)}, " ")
}
