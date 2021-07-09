package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type SetMysqlQuotasRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`

	Body *SetQuotasRequestBody `json:"body,omitempty"`
}

func (o SetMysqlQuotasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SetMysqlQuotasRequest struct{}"
	}

	return strings.Join([]string{"SetMysqlQuotasRequest", string(data)}, " ")
}
