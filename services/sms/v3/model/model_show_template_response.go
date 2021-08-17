package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowTemplateResponse struct {
	Template       *TemplateResponseBody `json:"template,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowTemplateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateResponse", string(data)}, " ")
}
