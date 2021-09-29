package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListGaussMySqlConfigurationsRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListGaussMySqlConfigurationsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListGaussMySqlConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlConfigurationsRequest", string(data)}, " ")
}
