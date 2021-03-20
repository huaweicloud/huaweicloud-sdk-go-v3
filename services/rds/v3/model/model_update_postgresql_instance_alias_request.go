package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePostgresqlInstanceAliasRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *UpdateRdsInstanceAliasRequest `json:"body,omitempty"`
}

func (o UpdatePostgresqlInstanceAliasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePostgresqlInstanceAliasRequest struct{}"
	}

	return strings.Join([]string{"UpdatePostgresqlInstanceAliasRequest", string(data)}, " ")
}
