package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateResourceResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateResourceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateResourceResponse struct{}"
	}

	return strings.Join([]string{"CreateResourceResponse", string(data)}, " ")
}
