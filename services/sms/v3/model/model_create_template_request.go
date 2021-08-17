package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateTemplateRequest struct {
	Body *CreateTemplateReq `json:"body,omitempty"`
}

func (o CreateTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateTemplateRequest", string(data)}, " ")
}
