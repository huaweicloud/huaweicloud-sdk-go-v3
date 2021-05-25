package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RestoreTablesRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *RestoreTablesRequestBody `json:"body,omitempty"`
}

func (o RestoreTablesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreTablesRequest struct{}"
	}

	return strings.Join([]string{"RestoreTablesRequest", string(data)}, " ")
}
