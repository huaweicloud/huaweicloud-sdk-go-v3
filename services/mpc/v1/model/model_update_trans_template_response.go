package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateTransTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTransTemplateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTransTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateTransTemplateResponse", string(data)}, " ")
}
