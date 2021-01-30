package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RestoreInstanceRequest struct {
	XLanguage *string                     `json:"X-Language,omitempty"`
	Body      *RestoreInstanceRequestBody `json:"body,omitempty"`
}

func (o RestoreInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreInstanceRequest", string(data)}, " ")
}
