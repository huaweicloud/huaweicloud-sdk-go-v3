package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListPostgresqlDatabasesRequest struct {
	XLanguage  *string `json:"X-Language,omitempty"`
	InstanceId string  `json:"instance_id"`
	Page       int32   `json:"page"`
	Limit      int32   `json:"limit"`
}

func (o ListPostgresqlDatabasesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPostgresqlDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListPostgresqlDatabasesRequest", string(data)}, " ")
}
