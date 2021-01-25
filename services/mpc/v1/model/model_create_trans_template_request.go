package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateTransTemplateRequest struct {
	Body *TransTemplate `json:"body,omitempty"`
}

func (o CreateTransTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTransTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateTransTemplateRequest", string(data)}, " ")
}
