package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListMysqlConfigurationsRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListMysqlConfigurationsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListMysqlConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"ListMysqlConfigurationsRequest", string(data)}, " ")
}
