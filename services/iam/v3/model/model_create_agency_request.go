package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAgencyRequest struct {
	Body *CreateAgencyRequestBody `json:"body,omitempty"`
}

func (o CreateAgencyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAgencyRequest struct{}"
	}

	return strings.Join([]string{"CreateAgencyRequest", string(data)}, " ")
}
