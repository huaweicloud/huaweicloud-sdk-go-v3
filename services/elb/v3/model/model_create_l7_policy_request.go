package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateL7PolicyRequest struct {
	Body *CreateL7PolicyRequestBody `json:"body,omitempty"`
}

func (o CreateL7PolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateL7PolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateL7PolicyRequest", string(data)}, " ")
}
