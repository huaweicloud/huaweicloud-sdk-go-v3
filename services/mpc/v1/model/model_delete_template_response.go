package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTemplateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplateResponse", string(data)}, " ")
}
