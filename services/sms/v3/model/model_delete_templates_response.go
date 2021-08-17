package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteTemplatesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTemplatesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTemplatesResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplatesResponse", string(data)}, " ")
}
