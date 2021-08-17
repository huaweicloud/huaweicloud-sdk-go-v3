package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTemplateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateTemplateResponse", string(data)}, " ")
}
