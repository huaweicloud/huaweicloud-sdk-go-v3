package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSqlserverDbUsersRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Page int32 `json:"page"`

	Limit int32 `json:"limit"`
}

func (o ListSqlserverDbUsersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSqlserverDbUsersRequest struct{}"
	}

	return strings.Join([]string{"ListSqlserverDbUsersRequest", string(data)}, " ")
}
