package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateTransTemplateRequest struct {
	Body *ModifyTransTemplateReq `json:"body,omitempty"`
}

func (o UpdateTransTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTransTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateTransTemplateRequest", string(data)}, " ")
}
