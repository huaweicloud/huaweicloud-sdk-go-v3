package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateCustomLineRequest struct {
	Body *CreateCustomLines `json:"body,omitempty"`
}

func (o CreateCustomLineRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateCustomLineRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomLineRequest", string(data)}, " ")
}
