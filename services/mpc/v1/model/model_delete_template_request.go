package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteTemplateRequest struct {
	TemplateId int32 `json:"template_id"`
}

func (o DeleteTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteTemplateRequest", string(data)}, " ")
}
