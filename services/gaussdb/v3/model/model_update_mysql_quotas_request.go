package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateMysqlQuotasRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`

	Body *SetQuotasRequestBody `json:"body,omitempty"`
}

func (o UpdateMysqlQuotasRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateMysqlQuotasRequest struct{}"
	}

	return strings.Join([]string{"UpdateMysqlQuotasRequest", string(data)}, " ")
}
