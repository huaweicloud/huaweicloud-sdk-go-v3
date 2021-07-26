package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateHostRequest struct {
	Body *CreateHostRequestBody `json:"body,omitempty"`
}

func (o CreateHostRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateHostRequest struct{}"
	}

	return strings.Join([]string{"CreateHostRequest", string(data)}, " ")
}
